import os
import sys

backend_path = os.path.abspath(os.path.join(os.getcwd(), 'backend'))
src_path = os.path.abspath(os.path.join(backend_path, 'src'))

if src_path not in sys.path:
    sys.path.append(src_path)

from backend.src import create_app
app = create_app()

if __name__ == "__main__":
    app.run(host='0.0.0.0', port='8000')