import os
import re
import hashlib
import requests
from urllib.parse import urlparse

CSS_URL_REGEX = re.compile(r'url\((["\']?)(https?://[^)\'"]+)\1\)')

def safe_filename(url: str) -> str:
    """
    Generate a stable filename from URL
    """
    parsed = urlparse(url)
    filename = os.path.basename(parsed.path)

    if filename:
        return filename

    # fallback (rare)
    h = hashlib.sha256(url.encode()).hexdigest()[:10]
    return f"asset-{h}"

def download_file(url: str, output_dir: str) -> str:
    os.makedirs(output_dir, exist_ok=True)

    filename = safe_filename(url)
    local_path = os.path.join(output_dir, filename)

    if os.path.exists(local_path):
        print(f"✔ Already exists: {filename}")
        return filename

    try:
        print(f"⬇ Downloading: {url}")
        response = requests.get(url, timeout=30)
        response.raise_for_status()

        with open(local_path, "wb") as f:
            f.write(response.content)

        return filename

    except requests.RequestException as e:
        print(f"❌ Failed to download {url}")
        raise RuntimeError(e)

def process_css(input_css: str, output_css: str, assets_dir: str):
    if not os.path.isfile(input_css):
        raise FileNotFoundError(f"CSS file not found: {input_css}")

    with open(input_css, "r", encoding="utf-8") as f:
        css_content = f.read()

    replacements = {}

    for match in CSS_URL_REGEX.finditer(css_content):
        remote_url = match.group(2)

        if remote_url in replacements:
            continue

        filename = download_file(remote_url, assets_dir)

        # Build relative path for CSS
        css_path = os.path.relpath(
            os.path.join(assets_dir, filename),
            start=os.path.dirname(output_css) or "."
        )

        replacements[remote_url] = css_path.replace("\\", "/")

    for remote, local in replacements.items():
        css_content = css_content.replace(remote, local)

    with open(output_css, "w", encoding="utf-8") as f:
        f.write(css_content)

    print("\n✅ Done")
    print(f"📄 Output CSS: {output_css}")
    print(f"📁 Assets saved in: {assets_dir}")

def ask_user_inputs():
    print("=== CSS Asset Downloader ===\n")

    input_css = input("Input CSS file path: ").strip()
    output_css = input("Output CSS file path: ").strip()
    assets_dir = input("Assets directory (e.g. static/fonts): ").strip()

    return input_css, output_css, assets_dir

if __name__ == "__main__":
    try:
        input_css, output_css, assets_dir = ask_user_inputs()
        process_css(input_css, output_css, assets_dir)
    except Exception as e:
        print(f"\n❌ Error: {e}")
