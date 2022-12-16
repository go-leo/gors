package com.leo.gors.intellij.handlers;

import com.intellij.codeInsight.editorActions.CommentCompleteHandler;
import com.intellij.lang.CodeDocumentationAwareCommenter;
import com.intellij.openapi.editor.Editor;
import com.intellij.psi.PsiComment;
import com.intellij.psi.PsiFile;

public class CommentCompleteHandlerImpl implements CommentCompleteHandler {
    @Override
    public boolean isCommentComplete(PsiComment comment, CodeDocumentationAwareCommenter commenter, Editor editor) {

        return false;
    }

    @Override
    public boolean isApplicable(PsiComment comment, CodeDocumentationAwareCommenter commenter) {
        return false;
    }
}
