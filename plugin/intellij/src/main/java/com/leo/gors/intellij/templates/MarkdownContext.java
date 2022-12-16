package com.leo.gors.intellij.templates;


import com.intellij.codeInsight.template.TemplateActionContext;
import com.intellij.codeInsight.template.TemplateContextType;
import org.jetbrains.annotations.NotNull;

public class MarkdownContext extends TemplateContextType {

    protected MarkdownContext() {
        super("MARKDOWN", "Markdown");
    }

    @Override
    public boolean isInContext(@NotNull TemplateActionContext templateActionContext) {
        return templateActionContext.getFile().getName().toLowerCase().endsWith(".md");
    }

}
