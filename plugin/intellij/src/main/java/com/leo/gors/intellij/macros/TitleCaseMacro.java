package com.leo.gors.intellij.macros;


import com.intellij.codeInsight.template.*;
import com.intellij.codeInsight.template.macro.MacroBase;
import com.intellij.openapi.util.text.StringUtil;
import com.leo.gors.intellij.templates.MarkdownContext;
import org.jetbrains.annotations.NotNull;

public class TitleCaseMacro extends MacroBase {

    public TitleCaseMacro() {
        this("titleCase", "titleCase(String)");
    }

    /**
     * Strictly to uphold contract for constructors in base class.
     */
    private TitleCaseMacro(String name, String description) {
        super(name, description);
    }

    @Override
    protected Result calculateResult(Expression @NotNull [] params, ExpressionContext context, boolean quick) {
        // Retrieve the text from the macro or selection, if any is available.
        String text = getTextResult(params, context, true);
        if (text == null) {
            return null;
        }
        if (text.length() > 0) {
            // Capitalize the start of every word
            text = StringUtil.toTitleCase(text);
        }
        return new TextResult(text);
    }

    @Override
    public boolean isAcceptableInContext(TemplateContextType context) {
        // Might want to be less restrictive in future
        return (context instanceof MarkdownContext);
    }

}
