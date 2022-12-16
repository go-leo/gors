package com.leo.gors.intellij.errors;

import com.intellij.openapi.diagnostic.ErrorReportSubmitter;
import com.intellij.openapi.util.NlsActions;
import org.jetbrains.annotations.NotNull;

public class SentryErrorReporter extends ErrorReportSubmitter {
    @Override
    public @NlsActions.ActionText @NotNull String getReportActionText() {

        return null;
    }
}
