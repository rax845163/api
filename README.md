# api
Alameda API definition

## How to generate the grpc files

### Go
1. Run the command
    ```bash
    for pt in `find . | grep \\\.proto$ | grep -v ^\\\./include`; do protoc -I . -I include/ $pt --go_out=plugins=grpc:.; done
    ```
### Python
1. Install the packages
    ```bash
    pip install -r requirements.txt
    ```
2. Run the command
    ```bash
    for pt in `find . | grep \\\.proto$ | grep -v ^\\\./include`; \
        do
            python -m grpc_tools.protoc -I . -I include/ --python_out=python/containers-ai_api --grpc_python_out=python/containers-ai_api $pt;
        done
    ```
