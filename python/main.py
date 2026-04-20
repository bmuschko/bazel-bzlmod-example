import requests


def get_languages():
    return ["Go", "Java", "Python"]


def fetch_json(url):
    response = requests.get(url)
    return {"status": response.status_code, "data": response.json()}


def main():
    print(f"Languages: {get_languages()}")
    result = fetch_json("https://httpbin.org/json")
    print(f"Status: {result['status']}")
    print(f"Response: {result['data']}")


if __name__ == "__main__":
    main()
