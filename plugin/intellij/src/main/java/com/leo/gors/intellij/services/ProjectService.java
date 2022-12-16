package com.leo.gors.intellij.services;

import com.intellij.openapi.project.Project;

public class ProjectService {

    private final Project myProject;

    public ProjectService(Project project) {
        myProject = project;
    }

    public void someServiceMethod(String parameter) {
        AnotherService anotherService = myProject.getService(AnotherService.class);
        String result = anotherService.anotherServiceMethod(parameter, false);
        // do some more stuff
    }
}
