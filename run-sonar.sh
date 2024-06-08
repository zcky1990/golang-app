docker run --rm --platform linux/amd64 -e SONAR_HOST_URL="http://192.168.1.9:9000" \
    -v "$(pwd):/usr/src" \
    sonarsource/sonar-scanner-cli \
    -Dsonar.projectBaseDir=/usr/src \
-Dsonar.token=sqp_7172f8abc9aa13b23b6314bdb2cd0bb721cc23e1