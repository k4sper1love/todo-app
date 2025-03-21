# ToDo App (Список задач) ☑️

Простое и стильное кроссплатформенное десктопное приложение для управления списком задач, разработанное с использованием фреймворка **Wails** (Go + Vue.js + Bootstrap). Позволяет добавлять, редактировать, удалять и отмечать задачи как выполненные. Состояние задач сохраняется в **SQLite** базе данных. Реализован **Repository Pattern** на бекенде и сервисный слой для взаимодействия с API на фронтенде

## ⭐ Основные функции
- **Добавление задач:**  Создание новых задач с текстом, приоритетом и опциональной датой выполнения.

- **Отметка выполнения:** Задачи можно отмечать как выполненные (с зачеркиванием текста и датой выполнения).

- **Удаление задач:** Удаление задач из списка с подтверждением.

- **Редактирование задач:** Изменение текста и даты выполнения задачи.

- **Приоритет задач:** Возможность пометить задачу как важную (с иконкой звезды).

- **Фильтрация задач:** Фильтрация задач по статусу (выполненные/невыполненные) и разбиение на разделы.

- **Сортировка задач:** Сортировка задач по приоритету, статусу, дате выполнения и дате создания.

- **Мониторинг дедлайнов:** Время выполнения задачи помечается красным цветом при истечении даты и времени выполнения.

- **Валидация данных:** Проверка введеных данных для избежания вставки пустого текста

- **Сохранение состояния:** Все задачи сохраняются в SQLite базе данных.

- **Стильный интерфейс:** Использование Bootstrap для красивого и отзывчивого дизайна.

## 🛠️ Стэк технологий
- **Backend:** Go (с использованием SQLite и Repository Pattern)
- **Frontend:** Vue.js + Bootstrap
- **Фреймворк:** Wails
- **База данных:** SQLite
- **Архитектура:**
  - **Repository Pattern** на бекенде для работы с базой данных.
  - Сервисный слой на фронтенде для взаимодействия с API.

## 📦 Установка и запуск
### Предварительные требования
Убедитесь, что установлены следующие инструменты:
1. **Go** (1.20+):
- Скачайте и установите с [оффициального сайта](https://go.dev/doc/install).
- Проверьте установку:
    ```bash
    go version
    ```
2. **Node.js** (15+):
- Скачайте и установите с [оффициального сайта](https://nodejs.org/en).
- Проверьте установку:
    ```bash
    node -v
    npm -v
    ```
2. **Wails** (2.0+):
- Установите Wails :
    ```bash
    go install github.com/wailsapp/wails/v2/cmd/wails@latest
    ```
- Проверьте установку:
    ```bash
    wails version
    ```
### Установка и запуск проекта
1. Клонируйте репозиторий:
```bash
git clone https://github.com/k4sper1love/todo-app.git
```
2. Перейдите в рабочую директорию проекта:
```bash
cd ваш-репозиторий/todoapp
```
3. Установите зависимости:
```bash
npm install
```
4. Запустите приложение в режиме разработки:
```bash
wails dev
```
5. Для сборки production-версии:
```bash
wails build
```

## 🖼️ Скриншоты
При первом запуске, запрашивает у пользователя имя

![image](https://github.com/user-attachments/assets/2066364b-f592-4d86-b376-3c92d3adde72)

Переходим в главный компонент приложение - список задач. Сверху приветствие, оно зависит от времени суток. Если нет активных задач - то выводится текст об этом. Если нет выполненных задач - мы не видим раздел с ними.

![image](https://github.com/user-attachments/assets/a711fcde-70fc-4644-9de3-0b732c22bcc1)

Добавление новой задачи, с возможность установки приоритета (иконка звезды) и выбора даты и времени дедлайна с помощью календаря

![image](https://github.com/user-attachments/assets/7f0463eb-3b5f-4a6e-bd05-5e345c915fc7)

После нажатия кнопки добавления задачи, она появлятся в списке активных задач

![image](https://github.com/user-attachments/assets/846d22a2-929c-489c-bc9e-cbfc57068a95)

При истечении дедлайна, время станет красным и отобразится выше задач, у которых время выполнения не истекло

![image](https://github.com/user-attachments/assets/b20e6d0b-879d-4b59-b878-835af8ef24ee)

При нажатии на иконку звезды, мы ставим задаче важный приоритет и она отображается выше задач без приоритета, либо наоборот, убираем приоритет

![image](https://github.com/user-attachments/assets/7e9ab68c-026b-4fdb-9d6f-c60aabae65b9)

При нажатии на текст задачи, мы либо делаем ее выполненной либо наоборот. У выполненной задачи зеленым отображается время выполнения и она становится перечеркнутой. При повторном нажатии, задача вернется в раздел активных и с установленным для нее дедлайном.

![image](https://github.com/user-attachments/assets/48719b80-d662-4880-ab4a-f66815f74ee0)

Мы можем скрыть раздел выполненных задач, нажав на треугольник

![image](https://github.com/user-attachments/assets/9138469c-b522-427f-af75-32a7a4fa49ac)

При нажатии на задачу (за исключением текста и кнопок приоритета, удаления), открывается модульное окно, где мы можем отредактировать задачу, изменив текст или дедлайн

![image](https://github.com/user-attachments/assets/3939c5fa-748e-4568-84dc-2c57e328392c)

При нажатии на кнопку удаления задачи, мы увидим модульное окно, которое потребует подтверждение действия

![image](https://github.com/user-attachments/assets/4aa00d65-70ba-472a-979a-d5fcbb08f224)

## 📜 Лицензия
Этот проект распространяется под [лицензией MIT](LICENSE.txt).

## 📫 Контакты
По любым вопросом или для обратной связи, пожалуйста свяжитесь с помощью:
- **Email**: s_yelkin@proton.me
- **Telegram**: [k4sper1love](https://t.me/k4sper1love)
- **GitHub**: [k4sper1love](https://github.com/k4sper1love)
