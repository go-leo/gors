package com.leo.gors.intellij.intentions;

import com.goide.intentions.GoBaseIntentionAction;
import com.goide.psi.*;
import com.goide.refactor.template.GoTemplate;
import com.intellij.codeInspection.util.IntentionFamilyName;
import com.intellij.codeInspection.util.IntentionName;
import com.intellij.openapi.application.ApplicationManager;
import com.intellij.openapi.application.WriteAction;
import com.intellij.openapi.diagnostic.Logger;
import com.intellij.openapi.editor.*;
import com.intellij.openapi.project.Project;
import com.intellij.openapi.ui.ComboBox;
import com.intellij.openapi.ui.DialogBuilder;
import com.intellij.openapi.util.TextRange;
import com.intellij.psi.*;
import com.intellij.psi.util.PsiTreeUtil;
import com.intellij.ui.CheckBoxList;
import com.intellij.ui.components.JBCheckBox;
import com.intellij.ui.components.JBLabel;
import com.intellij.ui.components.JBPanel;
import com.intellij.ui.components.JBTextField;
import com.intellij.ui.components.panels.HorizontalLayout;
import com.intellij.ui.components.panels.VerticalLayout;
import com.intellij.util.IncorrectOperationException;
import org.jdesktop.swingx.combobox.ListComboBoxModel;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import javax.swing.*;
import java.awt.*;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class GORSMethodIntentionAction extends GoBaseIntentionAction {
    private static final Logger LOG = Logger.getInstance(GORSMethodIntentionAction.class);

    public GORSMethodIntentionAction() {
    }

    @Override
    public @IntentionName @NotNull String getText() {
        return "Generate go restful method directive";
    }

    @Override
    public @NotNull @IntentionFamilyName String getFamilyName() {
        return "GORSMethodIntentionAction";
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
        final CaretModel caretModel = editor.getCaretModel();
        final Caret primaryCaret = caretModel.getPrimaryCaret();
        LogicalPosition logicalPosition = primaryCaret.getLogicalPosition();
        Document document = editor.getDocument();
        for (int i = 1; ; i++) {
            int lineStartOffset = document.getLineStartOffset(logicalPosition.line - i);
            int lineEndOffset = document.getLineEndOffset(logicalPosition.line - i);
            GoMethodSpec methodSpec = findMethodSpec(element);
            String text = document.getText(new TextRange(lineStartOffset, lineEndOffset));
            LOG.warn("pre: " + text);
            text = text.trim();
            LOG.warn("pre: " + text);
            if (text.startsWith("//")) {
                if (text.toUpperCase().contains("@GORS")) {
                    return false;
                } else {
                    continue;
                }
            } else {
                return true;
            }
        }
    }

    private static @Nullable GoMethodSpec findMethodSpec(@NotNull PsiElement element) {
        return PsiTreeUtil.getNonStrictParentOfType(element, GoMethodSpec.class);
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


        DialogBuilder dialogBuilder = new DialogBuilder();
        dialogBuilder.addOkAction();
        dialogBuilder.setTitle("Input Go Restful Service Method Params");

        JBPanel<JBPanel> jbPanel = new JBPanel<>(new VerticalLayout(10, SwingConstants.LEFT));

        JBPanel<JBPanel> methodPanel = new JBPanel<>(new HorizontalLayout(10));
        methodPanel.add(new JBLabel("Method:"));
        List<String> methods = Arrays.asList("@GET", "@POST", "@PUT", "@DELETE", "@PATCH", "@HEAD", "@CONNECT", "@OPTIONS", "@TRACE");
        ComboBox<String> methodComboBox = new ComboBox<String>(new ListComboBoxModel<>(methods));
        methodPanel.add(methodComboBox);
        jbPanel.add(methodPanel);

        JBPanel<JBPanel> pathPanel = new JBPanel<>(new HorizontalLayout(10));
        pathPanel.add(new JBLabel("Path:"));
        JBTextField pathText = new JBTextField(24);
        pathPanel.add(pathText);
        jbPanel.add(pathPanel);


        JBPanel<JBPanel> cbPanel = new JBPanel<>(new HorizontalLayout(10));
        cbPanel.add(new JBLabel("Common binding:"));
        ArrayList<JBCheckBox> jbCheckBoxes = new ArrayList<>();
        JBCheckBox uriBindingCB = new JBCheckBox("@UriBinding");
        cbPanel.add(uriBindingCB);
        jbCheckBoxes.add(uriBindingCB);
        JBCheckBox queryBindingCB = new JBCheckBox("@QueryBinding");
        cbPanel.add(queryBindingCB);
        jbCheckBoxes.add(queryBindingCB);
        JBCheckBox headerBindingCB = new JBCheckBox("@HeaderBinding");
        cbPanel.add(headerBindingCB);
        jbCheckBoxes.add(headerBindingCB);
        jbPanel.add(cbPanel);


        JBPanel<JBPanel> bbPanel = new JBPanel<>(new HorizontalLayout(10));
        bbPanel.add(new JBLabel("Body binding:"));
        List<String> bodyBindings = Arrays.asList("", "@JSONBinding", "@XMLBinding", "@FormBinding", "@FormPostBinding",
                "@FormMultipartBinding", "@ProtoBufBinding", "@MsgPackBinding", "@YAMLBinding", "@TOMLBinding");
        ComboBox<String> bodyBindingsComboBox = new ComboBox<String>(new ListComboBoxModel<>(bodyBindings));
        bbPanel.add(bodyBindingsComboBox);
        jbPanel.add(bbPanel);

        JBPanel<JBPanel> renderPanel = new JBPanel<>(new HorizontalLayout(10));
        renderPanel.add(new JBLabel("Render:"));
        List<String> renders = Arrays.asList(
                "@JSONRender",
                "@IndentedJSONRender",
                "@SecureJSONRender",
                "@JsonpJSONRender",
                "@PureJSONRender",
                "@AsciiJSONRender",
                "@XMLRender",
                "@YAMLRender",
                "@ProtoBufRender",
                "@MsgPackRender",
                "@TOMLRender",
                "@TextRender",
                "@HTMLRender",
                "@RedirectRender",
                "@BytesRender()",
                "@StringRender()",
                "@ReaderRender()");
        List<String> contentTypes = Arrays.asList(
                "application/json; charset=utf-8",
                "application/json; charset=utf-8",
                "application/json; charset=utf-8",
                "application/javascript; charset=utf-8",
                "application/json; charset=utf-8",
                "application/json",
                "application/xml; charset=utf-8",
                "application/x-yaml; charset=utf-8",
                "application/x-protobuf",
                "application/msgpack; charset=utf-8",
                "application/toml; charset=utf-8",
                "text/plain; charset=utf-8",
                "text/html; charset=utf-8",
                ""
        );
        ComboBox<String> rendersComboBox = new ComboBox<String>(new ListComboBoxModel<>(renders));
        renderPanel.add(rendersComboBox);
        JBTextField ctText = new JBTextField(22);
        ctText.setText(contentTypes.get(0));
        ctText.setEnabled(false);
        renderPanel.add(ctText);
        jbPanel.add(renderPanel);
        rendersComboBox.addItemListener(e -> {
            int selectedIndex = rendersComboBox.getSelectedIndex();
            String s = renders.get(selectedIndex);
            if (s.contains("()")) {
                ctText.setEnabled(true);
                ctText.setText("");
            } else {
                ctText.setEnabled(false);
                ctText.setText(contentTypes.get(selectedIndex));
            }
        });

        dialogBuilder.setCenterPanel(jbPanel);
        boolean showResult = dialogBuilder.showAndGet();
        if (!showResult) {
            return;
        }

        StringBuilder directive = new StringBuilder();
        directive.append("// @GORS ");
        directive.append(methods.get(methodComboBox.getSelectedIndex()));
        directive.append(" @Path(");
        final String path = pathText.getText();
        directive.append(path);
        directive.append(")");

        jbCheckBoxes.forEach(cb -> {
            if (!cb.isSelected()) {
                return;
            }
            directive.append(" ");
            directive.append(cb.getText());
        });

        int bodyBindingsSelectedIndex = bodyBindingsComboBox.getSelectedIndex();
        directive.append(" ");
        directive.append(bodyBindings.get(bodyBindingsSelectedIndex));

        int rendersSelectedIndex = rendersComboBox.getSelectedIndex();
        directive.append(" ");
        String str = renders.get(rendersSelectedIndex);
        if (str.contains("()")) {
            directive.append(str.substring(0, str.length() - 2));
            directive.append("(");
            directive.append(ctText.getText());
            directive.append(")");
        } else {
            directive.append(str);
        }

        directive.append("\n");

        GoMethodSpec methodSpec = findMethodSpec(element);

        WriteAction.runAndWait(() -> {
            LOG.warn("generate1: " + (!methodSpec.isValid()));
            if (!methodSpec.isValid()) {
                return;
            }
            String methodName = methodSpec.getName();
            GoTemplate template = new GoTemplate(methodSpec.getContainingFile());
            template.addTextSegment("// " + methodName + "\n");
            template.addTextSegment(directive.toString());
            int startOffset = methodSpec.getTextRange().getStartOffset();
            template.startTemplate(editor, startOffset, "generate gors service directive", null);
            editor.getCaretModel().getPrimaryCaret().moveToOffset(startOffset - 2);
        });

        WriteAction.run(() -> {

        });

    }

}

