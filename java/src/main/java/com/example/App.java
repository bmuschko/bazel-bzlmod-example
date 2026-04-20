package com.example;

import com.google.common.collect.ImmutableList;

public class App {
    public static void main(String[] args) {
        ImmutableList<String> languages = ImmutableList.of("Go", "Java", "Python");
        System.out.println("Languages in this Bazel workspace:");
        languages.forEach(lang -> System.out.println("  - " + lang));
    }
}
