package com.leo.gors.intellij.intentions;

import com.goide.intentions.GoBaseIntentionAction;
import com.goide.psi.*;
import com.goide.psi.impl.GoPsiUtil;
import com.goide.refactor.template.GoTemplate;
import com.intellij.codeInspection.util.IntentionFamilyName;
import com.intellij.codeInspection.util.IntentionName;
import com.intellij.openapi.application.ApplicationManager;
import com.intellij.openapi.application.WriteAction;
import com.intellij.openapi.diagnostic.Logger;
import com.intellij.openapi.editor.Editor;
import com.intellij.openapi.project.Project;
import com.intellij.openapi.ui.Messages;
import com.intellij.psi.PsiElement;
import com.intellij.psi.PsiFile;
import com.intellij.psi.util.PsiTreeUtil;
import com.intellij.util.IncorrectOperationException;
import com.intellij.util.ObjectUtils;
import com.leo.gors.intellij.actions.DialogAction;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;

public class GORSServiceIntentionAction extends GoBaseIntentionAction {
    private static final Logger LOG = Logger.getInstance(GORSServiceIntentionAction.class);

    public GORSServiceIntentionAction() {
    }

    @Override
    public @IntentionName @NotNull String getText() {
        return "Generate go restful service directive";
    }

    @Override
    public @NotNull @IntentionFamilyName String getFamilyName() {
        return "GORSServiceIntentionAction";
    }

    @Override
    public boolean startInWriteAction() {
        return false;
    }

    @Override
    public @Nullable PsiElement getElementToMakeWritable(@NotNull PsiFile currentFile) {
        return super.getElementToMakeWritable(currentFile);
    }

    @Override
    public boolean isAvailable(@NotNull Project project, Editor editor, @NotNull PsiElement element) {
        LOG.warn("1: " + (editor == null));
        if (editor == null) {
            return false;
        }
        LOG.warn("2: " + ApplicationManager.getApplication().isUnitTestMode());
        if (ApplicationManager.getApplication().isUnitTestMode()) {
            return false;
        }
        String text = editor.getDocument().getText();
        GoTypeSpec typeSpec = findTypeSpec(element);
        LOG.warn("3: " + !GoPsiUtil.isTopLevelDeclaration(typeSpec));
        if (!GoPsiUtil.isTopLevelDeclaration(typeSpec)) {
            return false;
        }
        LOG.warn("4: " + !(typeSpec.getSpecType().getType() instanceof GoInterfaceType));
        if (!(typeSpec.getSpecType().getType() instanceof GoInterfaceType)) {
            return false;
        }
        String typeName = typeSpec.getName();
        LOG.warn("5: " + text.contains("//go:generate gors -service " + typeName));
        if (text.contains("//go:generate gors -service " + typeName)) {
            return false;
        }
        return true;
    }

    private static @Nullable GoTypeSpec findTypeSpec(@NotNull PsiElement element) {
        PsiElement parent = PsiTreeUtil.getNonStrictParentOfType(element, GoTypeSpec.class, GoTypeDeclaration.class);
        if (parent instanceof GoTypeSpec) {
            return (GoTypeSpec) parent;
        }
        return null;
    }

    @Override
    public void invoke(@NotNull Project project, Editor editor, @NotNull PsiElement element) throws IncorrectOperationException {
        LOG.warn("invoke1: " + (editor == null));
        if (editor == null) {
            return;
        }
        LOG.warn("invoke2: " + (ApplicationManager.getApplication().isUnitTestMode()));
        if (ApplicationManager.getApplication().isUnitTestMode()) {
            return;
        }
        GoTypeSpec typeSpec = findTypeSpec(element);
        LOG.warn("invoke3: " + (!GoPsiUtil.isTopLevelDeclaration(typeSpec)));
        if (!GoPsiUtil.isTopLevelDeclaration(typeSpec)) {
            return;
        }
        GoInterfaceType interfaceType = ObjectUtils.tryCast(typeSpec.getSpecType().getType(), GoInterfaceType.class);
        LOG.warn("invoke4: " + (interfaceType == null));
        if (interfaceType == null) {
            return;
        }

        String path = Messages.showInputDialog(project,"Path:","Input Base Path", null);

        WriteAction.runAndWait(() -> {
            LOG.warn("generate1: " + (!typeSpec.isValid()));
            if (!typeSpec.isValid()) {
                return;
            }
            String typeName = typeSpec.getName();
            GoTypeDeclaration typeDeclaration = PsiTreeUtil.getParentOfType(typeSpec, GoTypeDeclaration.class);
            LOG.warn("generate2: " + (typeDeclaration == null));
            if (typeDeclaration == null) {
                return;
            }
            GoTemplate template = new GoTemplate(typeSpec.getContainingFile());
            template.addTextSegment("//go:generate gors -service " + typeName + "\n");
            template.addTextSegment("\n");
            template.addTextSegment("// " + typeName + "\n");
            template.addTextSegment("// @GORS @Path("+path+")\n");
            int startOffset = typeDeclaration.getTextRange().getStartOffset();
            template.startTemplate(editor, startOffset, "generate gors service directive", null);
            editor.getCaretModel().getPrimaryCaret().moveToOffset(startOffset - 2);
        });
        WriteAction.run(() -> {
            GoTypeDeclaration typeDeclaration = PsiTreeUtil.getParentOfType(typeSpec, GoTypeDeclaration.class);
            LOG.warn("generate2: " + (typeDeclaration == null));
            if (typeDeclaration == null) {
                return;
            }
            int startOffset = typeDeclaration.getTextRange().getStartOffset();
            editor.getCaretModel().getPrimaryCaret().moveToOffset(startOffset - 2);
        });
    }
}

