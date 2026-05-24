<<<<<<< HEAD
# SpotiFLACNeo

**Neo-редакция** русской локализации [SpotiFLAC](https://github.com/spotbye/SpotiFLAC) — десктопное приложение для загрузки треков Spotify в настоящем качестве **FLAC** из Tidal, Qobuz и Amazon Music. Аккаунты стриминговых сервисов не требуются.

| | |
|---|---|
| **Оригинал** | [spotbye/SpotiFLAC](https://github.com/spotbye/SpotiFLAC) |
| **Этот проект** | [krwg/SpotiFLACNeo](https://github.com/krwg/SpotiFLACNeo) |
| **Стек** | Go (Wails v2) + React + TypeScript |
| **Платформы** | Windows · macOS · Linux |

![Windows](https://img.shields.io/badge/Windows-10%2B-0078D6?style=flat-square)
![macOS](https://img.shields.io/badge/macOS-10.13%2B-000000?style=flat-square)
![Linux](https://img.shields.io/badge/Linux-Any-FCC624?style=flat-square)
[![Релизы](https://img.shields.io/github/v/release/krwg/SpotiFLACNeo?style=flat-square&label=релиз)](https://github.com/krwg/SpotiFLACNeo/releases)

---

## Чем отличается от оригинала

- **Neo edition** — обновлённый брендинг и интерфейс на базе русской локализации.
- **Полностью русский интерфейс** — меню, настройки, уведомления, ошибки.
- **Обновлённый UI/UX** — стиль в духе Apple (glass-панели, мягкие тени, скругления), шрифт **Monocraft** с поддержкой кириллицы для заголовков.
- **Локальный сборщик** — скрипт `build.ps1` для сборки `.exe` на Windows без ручной настройки.
- **Дорожная карта** — планируется поддержка **Яндекс Музыки**, **VK Музыки**, других российских площадок и **Apple Music**.

---

## Скриншот

<img width="1152" height="689" alt="image" src="https://github.com/user-attachments/assets/7a0db452-90ba-4162-b80c-cb571059737e" />


---

## Быстрый старт

### Скачать готовую сборку

1. Откройте [Releases](https://github.com/krwg/SpotiFLACNeo/releases).
2. Скачайте `SpotiFLACNeo.exe` (Windows) или артефакт для вашей ОС.
3. Запустите. При первом запуске приложение может предложить установить **FFmpeg** (~30–40 МБ).

### Собрать из исходников (Windows)

**Требования:**

| Инструмент | Версия |
|------------|--------|
| [Go](https://go.dev/dl/) | 1.26+ (после установки перезапустите терминал) |
| [Node.js](https://nodejs.org/) | 20+ (рекомендуется 24 LTS) |
| [Wails v2](https://wails.io/docs/gettingstarted/installation) | latest |
| WebView2 | обычно уже есть в Windows 10/11 |

**Установка Go (если `go` не найден):**

```powershell
winget install GoLang.Go
# или скачайте MSI: https://go.dev/dl/
go version
```

**Сборка одной командой:**

```bat
git clone https://github.com/krwg/SpotiFLACNeo.git
cd SpotiFLACNeo
build.bat
```

> Если `.\build.ps1` блокируется политикой PowerShell, используйте **`build.bat`** — он обходит ограничение только для этой сборки.

Готовый файл: `build\bin\SpotiFLACNeo.exe`

**Дополнительные параметры:**

```bat
build.bat -Dev          :: режим разработки (wails dev)
build.bat -SkipFrontend :: не переустанавливать npm-зависимости
build.bat -Compress     :: сжать exe через UPX (если установлен)
```

**Ручная сборка:**

```powershell
cd frontend
npm install
npm run build
cd ..
wails build -platform windows/amd64
```

### macOS / Linux

См. [официальную документацию Wails](https://wails.io/docs/gettingstarted/installation) и workflow [`.github/workflows/build.yml`](.github/workflows/build.yml). Релизы для macOS (`.dmg`) и Linux (`.AppImage`) собираются автоматически при пуше тега `v*`.

---

## Как пользоваться

1. Скопируйте ссылку на **трек**, **альбом**, **плейлист** или **исполнителя** из Spotify (`open.spotify.com/...`).
2. Вставьте в поле на главной странице и нажмите **Получить**.
3. Выберите треки и нажмите загрузку — файлы сохранятся в папку из **Настроек**.
4. При ошибках метаданных попробуйте **VPN** (США, Великобритания, Германия, Нидерланды, Сингапур).

Также доступны:

- **Режим поиска** — поиск по названию без ссылки.
- **История** — последние запросы.
- **Инструменты** — анализ качества, конвертер, ресемплер, файловый менеджер.

---

## Настройки

| Раздел | Описание |
|--------|----------|
| Путь загрузки | Папка для FLAC-файлов |
| Источник | Tidal / Qobuz / Amazon / авто |
| Тема | Цвет акцента и светлая/тёмная тема |
| Шрифт | Apple System, Monocraft (Minecraft), Google Sans и др. |
| Шаблоны имён | Папки и имена файлов с плейсхолдерами |

---

## CI/CD

При создании тега `v*` (например `v1.0.0`) GitHub Actions:

1. Собирает приложение для Windows, macOS и Linux.
2. Создаёт **draft**-релиз с артефактами.

Локально Windows-сборку дублирует `build.ps1`.

---

## Часто задаваемые вопросы

<details>
<summary>Программа бесплатна?</summary>

Да. Регистрация и подписки не нужны — только интернет.
</details>

<details>
<summary>Забанят ли Spotify-аккаунт?</summary>

Нет. Приложение **не подключается** к вашему аккаунту Spotify. Метаданные получаются через публичный веб-плеер, без авторизации.
</details>

<details>
<summary>Откуда берётся аудио?</summary>

Из сторонних API (Tidal, Qobuz, Amazon Music и др.) по сопоставлению с метаданными Spotify.
</details>

<details>
<summary>Не загружаются метаданные</summary>

Часто это **rate limit по IP**. Подождите или используйте VPN в регионе без блокировки Spotify.
</details>

<details>
<summary>Антивирус удаляет .exe</summary>

Возможно ложное срабатывание (в CI используется сжатие UPX). Соберите сами через `build.ps1` или отключите UPX.
</details>

<details>
<summary>Когда Яндекс / VK / Apple Music?</summary>

В разработке. Следите за [Issues](https://github.com/krwg/SpotiFLACNeo/issues) и релизами.
</details>

---

## Поддержка проекта

- **SpotiFLACNeo:** [DonationAlerts](https://dalink.to/fallnix)
- **Оригинал:** [Ko-fi — afkarxyz](https://ko-fi.com/afkarxyz)

---

## Правовая информация

Проект предназначен **только для образовательных и личных целей**. Разработчики не поощряют нарушение авторских прав.

**SpotiFLACNeo** не связан со Spotify, Tidal, Qobuz, Amazon Music, Яндекс Музыкой, VK Музыкой, Apple Music и другими сервисами.

Вы несёте ответственность за соблюдение местного законодательства и условий использования платформ.

ПО предоставляется «как есть», без гарантий.

---

## Благодарности

- [spotbye / SpotiFLAC](https://github.com/spotbye/SpotiFLAC) — оригинальный проект
- [Wails](https://wails.io) — фреймворк десктопных приложений
- [Monocraft](https://github.com/IdreesInc/Monocraft) — шрифт в стиле Minecraft (OFL 1.1, кириллица)
- API: MusicBrainz · LRCLIB · Songlink/Odesli · и др. (см. оригинальный README)

---

## Лицензия

Исходный код распространяется по лицензии **MIT** (как у оригинала). Шрифт Monocraft — **SIL Open Font License 1.1**.

**Оригинал:** [afkarxyz / spotbye](https://github.com/spotbye/SpotiFLAC)  
**SpotiFLACNeo:** [krwg](https://github.com/krwg/SpotiFLACNeo)
=======
# SpotiFLAC

<a href="https://trendshift.io/repositories/15737" target="_blank"><img src="https://trendshift.io/api/badge/repositories/15737" alt="afkarxyz%2FSpotiFLAC | Trendshift" style="width: 250px; height: 55px;" width="250" height="55"/></a>

Get Spotify tracks in true FLAC from Tidal, Qobuz & Amazon Music — no account required.

![Windows](https://img.shields.io/badge/Windows-10%2B-0078D6?style=for-the-badge&logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI1MTIiIGhlaWdodD0iNTEyIiB2aWV3Qm94PSIwIDAgMjAgMjAiPjxwYXRoIGZpbGw9IiNmZmZmZmYiIGZpbGwtcnVsZT0iZXZlbm9kZCIgZD0iTTIwIDEwLjg3M1YyMEw4LjQ3OSAxOC41MzdsLjAwMS03LjY2NEgyMFptLTEzLjEyIDBsLS4wMDEgNy40NjFMMCAxNy40NjF2LTYuNTg4aDYuODhaTTIwIDkuMjczSDguNDhsLS4wMDEtNy44MUwyMCAwdjkuMjczWk02Ljg3OSAxLjY2NmwuMDAxIDcuNjA3SDBWMi41MzlsNi44NzktLjg3M1oiLz48L3N2Zz4=)
![macOS](https://img.shields.io/badge/macOS-10.13%2B-000000?style=for-the-badge&logo=apple&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-Any-FCC624?style=for-the-badge&logo=linux&logoColor=white)
[![Announcements](https://img.shields.io/badge/ANNOUNCEMENTS-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)](https://t.me/spotiflac)
[![Chat](https://img.shields.io/badge/CHAT-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)](https://t.me/spotiflac_chat)

### [Download](https://github.com/spotbye/SpotiFLAC/releases)

![Image](https://github.com/user-attachments/assets/c2624ca5-8569-49f0-950e-4410b523cea1)

## Other projects

### [SpotiFLAC Next](https://github.com/spotbye/SpotiFLAC-Next)

Get Spotify tracks in true Lossless from Tidal, Qobuz, Amazon Music, Deezer & Apple Music — no account required.

### [SpotubeDL.com](https://spotubedl.com)

Download Spotify Tracks, Albums, Playlists & Discography as MP3/OGG/Opus.

## Related projects

### [SpotiFLAC (Mobile)](https://github.com/zarzet/SpotiFLAC-Mobile)

SpotiFLAC for Android & iOS — maintained by [@zarzet](https://github.com/zarzet)

### [SpotiFLAC (Python Module)](https://github.com/ShuShuzinhuu/SpotiFLAC-Module-Version)

SpotiFLAC Python library for SpotiFLAC integration — maintained by [@ShuShuzinhuu](https://github.com/ShuShuzinhuu)

## FAQ

<details>
<summary>Is this software free?</summary>

_Yes. This software is completely free.
You do not need an account, login, or subscription.
All you need is an internet connection._

</details>

<details>
<summary>Can using this software get my Spotify account suspended or banned?</summary>

_No.
This software has no connection to your Spotify account.
Spotify data is obtained through reverse engineering of the Spotify Web Player, not through user authentication._

</details>

<details>
<summary>Where does the audio come from?</summary>

_The audio is fetched using third-party APIs._

</details>

<details>
<summary>Why does metadata fetching sometimes fail?</summary>

_This usually happens because your IP address has been rate-limited.
You can wait and try again later, or use a VPN to bypass the rate limit._

</details>

<details>
<summary>Why does Windows Defender or antivirus flag or delete the file?</summary>

_This is a false positive.
It likely happens because the executable is compressed using UPX._

_If you are concerned, you can fork the repository and build the software yourself from source._

</details>

<details>
<summary>Want to support the project?</summary>

_If this software is useful and brings you value,
consider supporting the project by buying me a coffee.
Your support helps keep development going._

</details>

[![Ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/afkarxyz)

## Disclaimer

This project is for **educational and private use only**. The developer does not condone or encourage copyright infringement.

**SpotiFLAC** is a third-party tool and is not affiliated with, endorsed by, or connected to Spotify, Tidal, Qobuz, Amazon Music or any other streaming service.

You are solely responsible for:

1. Ensuring your use of this software complies with your local laws.
2. Reading and adhering to the Terms of Service of the respective platforms.
3. Any legal consequences resulting from the misuse of this tool.

The software is provided "as is", without warranty of any kind. The author assumes no liability for any bans, damages, or legal issues arising from its use.

## API Credits

[MusicBrainz](https://musicbrainz.org) · [LRCLIB](https://lrclib.net) · [Songlink/Odesli](https://song.link) · [hifi-api](https://github.com/binimum/hifi-api) · [WJHE](https://music.wjhe.top) · [GDStudio](https://music.gdstudio.xyz) · [MusicDL](https://musicdl.me)

> [!TIP]
>
> **Star Us**, You will receive all release notifications from GitHub without any delay ~

[![RepoStars](https://repostars.dev/api/embed?repo=afkarxyz%2FSpotiFLAC&theme=forest)](https://repostars.dev/?repos=afkarxyz%2FSpotiFLAC&theme=forest)
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
