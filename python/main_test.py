import unittest
from unittest.mock import patch, Mock

from main import get_languages, fetch_json


class TestGetLanguages(unittest.TestCase):
    def test_returns_three_languages(self):
        self.assertEqual(len(get_languages()), 3)

    def test_contains_expected_languages(self):
        languages = get_languages()
        self.assertIn("Go", languages)
        self.assertIn("Java", languages)
        self.assertIn("Python", languages)


class TestFetchJson(unittest.TestCase):
    @patch("main.requests.get")
    def test_returns_status_and_data(self, mock_get):
        mock_response = Mock()
        mock_response.status_code = 200
        mock_response.json.return_value = {"key": "value"}
        mock_get.return_value = mock_response

        result = fetch_json("https://example.com/api")

        self.assertEqual(result["status"], 200)
        self.assertEqual(result["data"], {"key": "value"})
        mock_get.assert_called_once_with("https://example.com/api")


if __name__ == "__main__":
    unittest.main()
