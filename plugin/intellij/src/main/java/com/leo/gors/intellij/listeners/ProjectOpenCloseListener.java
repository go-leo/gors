package com.leo.gors.intellij.listeners;


import com.intellij.openapi.application.ApplicationManager;
import com.intellij.openapi.project.Project;
import com.intellij.openapi.project.ProjectManagerListener;
import com.intellij.openapi.ui.Messages;
import com.leo.gors.intellij.services.ProjectCountingService;
import org.jetbrains.annotations.NotNull;

/**
 * Listener to detect project open and close.
 * Depends on {@link ProjectCountingService}
 */
public class ProjectOpenCloseListener implements ProjectManagerListener {

    /**
     * Invoked on project open.
     *
     * @param project opening project
     */
    @Override
    public void projectOpened(@NotNull Project project) {
        // Ensure this isn't part of testing
        if (ApplicationManager.getApplication().isUnitTestMode()) {
            return;
        }
        // Get the counting service
        ProjectCountingService projectCountingService =
                ApplicationManager.getApplication().getService(ProjectCountingService.class);
        // Increment the project count
        projectCountingService.incrProjectCount();
        // See if the total # of projects violates the limit.
        if (projectCountingService.projectLimitExceeded()) {
            // Transitioned to outside the limit
            String title = String.format("Opening Project \"%s\"", project.getName());
            String message = "<br>The number of open projects exceeds the SDK plugin max_opened_projects limit.<br><br>" +
                    "This is not an error<br><br>";
            Messages.showMessageDialog(project, message, title, Messages.getInformationIcon());
        }
    }

    /**
     * Invoked on project close.
     *
     * @param project closing project
     */
    @Override
    public void projectClosed(@NotNull Project project) {
        // Ensure this isn't part of testing
        if (ApplicationManager.getApplication().isUnitTestMode()) {
            return;
        }
        // Get the counting service
        ProjectCountingService projectCountingService =
                ApplicationManager.getApplication().getService(ProjectCountingService.class);
        // Decrement the count because a project just closed
        projectCountingService.decrProjectCount();
    }

    @Override
    public void projectClosing(@NotNull Project project) {

    }

    @Override
    public void projectClosingBeforeSave(@NotNull Project project) {

    }
}