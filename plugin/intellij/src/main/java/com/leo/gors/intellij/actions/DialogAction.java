package com.leo.gors.intellij.actions;

import com.intellij.openapi.actionSystem.AnAction;
import com.intellij.openapi.actionSystem.AnActionEvent;
import com.intellij.openapi.diagnostic.Logger;
import com.leo.gors.intellij.dialogs.SampleDialogWrapper;
import org.jetbrains.annotations.NotNull;

/**
 * @author songyancheng
 */
public class DialogAction extends AnAction {
    private static final Logger LOG = Logger.getInstance(DialogAction.class);

    // The update() method implements the code that enables or disables an action.
    @Override
    public void update(@NotNull AnActionEvent event) {
        // Using the event, evaluate the context,
        // and enable or disable the action.
        super.update(event);
    }

    // The actionPerformed() method implements the code that executes when an action is invoked by the user.
    @Override
    public void actionPerformed(@NotNull AnActionEvent e) {
        LOG.info("hello PopupDialogAction");
        if (new SampleDialogWrapper().showAndGet()) {
            // user pressed OK
        }
    }
}
