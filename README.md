<a name="readme-top"></a>

<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <h3 align="center">WSL Path Manager</h3>

  <p align="center">
    Version: 0.0.2
    <br />
    An elegant desktop application to seamlessly manage WSL executable paths in Windows.
    <br />
    <a href="https://github.com/NextXierra/wslpathmanager"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/NextXierra/wslpathmanager">View Demo</a>
    ·
    <a href="https://github.com/NextXierra/wslpathmanager/issues">Report Bug</a>
    ·
    <a href="https://github.com/NextXierra/wslpathmanager/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

WSL Path Manager is a desktop application that allows users to seamlessly manage and inject Windows Subsystem for Linux (WSL) executable paths into the Windows environment variables. It provides a user interface to scan, select, and create Windows shims for WSL tools, enabling native-like execution of Linux commands directly from Windows.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

* [![Go][Go-shield]][Go-url]
* [![VanillaJS][VanillaJS-shield]][VanillaJS-url]
* [![Wails][Wails-shield]][Wails-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running, follow these simple steps.

### Prerequisites

* Windows Operating System with WSL installed and configured.
* Go 1.20 or higher.
* Wails CLI v2.

### Installation

1. Clone the repository
   ```sh
   git clone https://github.com/NextXierra/wslpathmanager.git
   ```
2. Navigate to the project directory
   ```sh
   cd wslpathmanager
   ```
3. Build the application using Wails
   ```sh
   wails build
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE -->
## Usage

1. Launch the compiled executable.
2. Select a target WSL distribution from the provided dropdown menu.
3. Click "Scan Paths" to discover available tools and executables within the selected WSL distribution.
4. Select the specific tools you wish to make available in the Windows environment.
5. Click "Save & Apply Settings" to generate the appropriate shims and automatically inject them into your Windows PATH.
6. Check the Settings page to toggle System Tray integration and Auto-Startup features.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->
## Roadmap

- [x] Initial release (v0.0.1)
- [x] Add support for custom tool paths (v0.0.2)
- [x] Implement system tray support (v0.0.2)
- [x] Implement run on system startup (v0.0.2)
- [ ] Multi-language support

See the [open issues](https://github.com/NextXierra/wslpathmanager/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

NextXierra - 2411016110002@mhs.ulm.ac.id

Project Link: [https://github.com/NextXierra/wslpathmanager](https://github.com/NextXierra/wslpathmanager)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [Wails Framework](https://wails.io/)
* [Go](https://go.dev/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
[contributors-shield]: https://img.shields.io/github/contributors/NextXierra/wslpathmanager.svg?style=for-the-badge
[contributors-url]: https://github.com/NextXierra/wslpathmanager/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/NextXierra/wslpathmanager.svg?style=for-the-badge
[forks-url]: https://github.com/NextXierra/wslpathmanager/network/members
[stars-shield]: https://img.shields.io/github/stars/NextXierra/wslpathmanager.svg?style=for-the-badge
[stars-url]: https://github.com/NextXierra/wslpathmanager/stargazers
[issues-shield]: https://img.shields.io/github/issues/NextXierra/wslpathmanager.svg?style=for-the-badge
[issues-url]: https://github.com/NextXierra/wslpathmanager/issues
[license-shield]: https://img.shields.io/github/license/NextXierra/wslpathmanager.svg?style=for-the-badge
[license-url]: https://github.com/NextXierra/wslpathmanager/blob/master/LICENSE.txt
[Go-shield]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://go.dev/
[VanillaJS-shield]: https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E
[VanillaJS-url]: https://developer.mozilla.org/en-US/docs/Web/JavaScript
[Wails-shield]: https://img.shields.io/badge/Wails-ED1D24?style=for-the-badge&logo=wails&logoColor=white
[Wails-url]: https://wails.io/
