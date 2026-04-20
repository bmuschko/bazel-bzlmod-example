package com.example;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertTrue;

import com.google.common.collect.ImmutableList;
import org.junit.Test;

public class AppTest {
    @Test
    public void testGetLanguagesSize() {
        ImmutableList<String> languages = App.getLanguages();
        assertEquals(3, languages.size());
    }

    @Test
    public void testGetLanguagesContents() {
        ImmutableList<String> languages = App.getLanguages();
        assertTrue(languages.contains("Go"));
        assertTrue(languages.contains("Java"));
        assertTrue(languages.contains("Python"));
    }
}
