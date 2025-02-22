# Используем базовый образ Ubuntu
FROM ubuntu:latest

# Устанавливаем необходимые зависимости
RUN apt-get update && apt-get install -y \
    dpkg \
    && rm -rf /var/lib/apt/lists/*

# Копируем .deb пакет в контейнер
COPY myprogram.deb /tmp/myprogram.deb

# Устанавливаем .deb пакет
RUN dpkg -i /tmp/myprogram.deb || apt-get install -f -y

# Указываем команду для запуска вашего ПО
CMD ["lab", "--success-param"]