import requests


def main():
    response = requests.get("https://httpbin.org/json")
    print(f"Status: {response.status_code}")
    print(f"Content-Type: {response.headers['Content-Type']}")
    data = response.json()
    print(f"Response: {data}")


if __name__ == "__main__":
    main()
