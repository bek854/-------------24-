# Todo List App with CI/CD

## Docker Deployment
Images are automatically built and pushed to DockerHub on git tags:

```bash
git tag v1.0.0
git push origin v1.0.0
# Создайте новую ветку с улучшениями
git checkout -b improve-ci-cd

# Добавим реальную проверку Docker образа
cat > .github/workflows/ci-cd.yml << 'EOF'
name: CI/CD Pipeline

on:
  push:
    branches: [ main ]
    tags: [ 'v*.*.*' ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    name: Run Code Checks
    runs-on: ubuntu-latest
    steps:
    - name: Checkout code
      uses: actions/checkout@v4
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.23'
    - name: Verify dependencies
      run: |
        cd todo-list
        go mod download
        go mod verify
    - name: Check code quality
      run: |
        cd todo-list
        go vet ./...
        go fmt ./...

  build-and-push:
    name: Build and Push Docker Image
    runs-on: ubuntu-latest
    needs: test
    if: startsWith(github.ref, 'refs/tags/v')
    steps:
    - name: Checkout code
      uses: actions/checkout@v4
    - name: Set up Docker Buildx
      uses: docker/setup-buildx-action@v2
    - name: Log in to DockerHub
      uses: docker/login-action@v2
      with:
        username: ${{ secrets.DOCKERHUB_USERNAME }}
        password: ${{ secrets.DOCKERHUB_TOKEN }}
    - name: Build and push
      uses: docker/build-push-action@v4
      with:
        context: ./todo-list
        push: true
        tags: |
          ${{ secrets.DOCKERHUB_USERNAME }}/todo-list-app:${{ github.ref_name }}
          ${{ secrets.DOCKERHUB_USERNAME }}/todo-list-app:latest
    - name: Verify pushed image
      run: |
        echo "Pulling image from DockerHub..."
        docker pull ${{ secrets.DOCKERHUB_USERNAME }}/todo-list-app:${{ github.ref_name }}
        echo "✅ Image successfully pushed to DockerHub:"
        echo "   - ${{ secrets.DOCKERHUB_USERNAME }}/todo-list-app:${{ github.ref_name }}"
        echo "   - ${{ secrets.DOCKERHUB_USERNAME }}/todo-list-app:latest"
