package com.leo.gors.intellij.actions;

import com.intellij.openapi.actionSystem.AnAction;
import com.intellij.openapi.actionSystem.AnActionEvent;
import com.intellij.openapi.actionSystem.CommonDataKeys;
import com.intellij.openapi.editor.actionSystem.EditorActionManager;
import com.intellij.openapi.editor.actionSystem.TypedAction;
import com.intellij.openapi.editor.actionSystem.TypedActionHandler;
import com.intellij.psi.PsiFile;
import com.leo.gors.intellij.handlers.TypedActionHandlerImpl;
import org.jetbrains.annotations.NotNull;

import java.util.concurrent.Callable;

@Deprecated
public class TypedActionHandlerAction extends AnAction {
    static {
        EditorActionManager actionManager = EditorActionManager.getInstance();
        TypedAction typedAction = actionManager.getTypedAction();
        typedAction.setupHandler(new TypedActionHandlerImpl());
    }

    @Override
    public void actionPerformed(@NotNull AnActionEvent e) {
    }
}
