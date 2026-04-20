package com.example;

import com.google.common.collect.ImmutableList;

public class App {
    public static ImmutableList<String> getLanguages() {
        return ImmutableList.of("Go", "Java", "Python");
    }

    public static void main(String[] args) {
        System.out.println("Languages in this Bazel workspace:");
        getLanguages().forEach(lang -> System.out.println("  - " + lang));
    }
}
