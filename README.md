# Raspberry Pi Dashboard

A simple dashboard to monitor the CPU, RAM, disk, and network usage of a Raspberry Pi.

## Features

- Monitor CPU usage and temperature
- Monitor RAM usage
- Monitor disk usage
- Monitor network usage
- Log important events, errors, and information

## Installation

1. Clone the repository:

```bash
git clone https://github.com/karmek-k/rpi-dashboard.git
```

2. Navigate to the project directory:

```bash
cd raspberry-pi-dashboard
```

3. Install the dependencies:

```bash
go mod tidy
```

## Usage

1. Run the server:

```bash
go run main.go
```

2. Access the dashboard at `http://localhost:8000` in your web browser.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the GNU General Public License v3.0. See the `LICENSE` file for more information.