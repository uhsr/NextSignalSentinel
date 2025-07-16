# NextSignalSentinel: Real-time Crypto Anomaly Detection

Detecting anomalous price movements in cryptocurrency markets with low-latency WebSocket alerts, powered by spectral analysis and a custom Kalman filter.

NextSignalSentinel provides a robust platform for identifying unusual price fluctuations in cryptocurrency markets in real-time. By leveraging sophisticated signal processing techniques, specifically spectral analysis using Fast Fourier Transforms (FFTs), the system pinpoints deviations from established price patterns. These deviations are then filtered through a custom-built Kalman filter, enhancing signal clarity and minimizing false positives. The result is a highly accurate and responsive anomaly detection system ideal for traders, arbitrageurs, and risk management professionals seeking a competitive edge in the volatile crypto landscape. Unlike simple moving average-based alerts, NextSignalSentinel identifies subtle but potentially significant changes in market behavior indicative of nascent trends or unforeseen events.

The core functionality centers around ingesting real-time price data from cryptocurrency exchanges via WebSocket connections. This raw data is then pre-processed, transformed into the frequency domain using FFTs, and analyzed for spectral anomalies. The Kalman filter plays a crucial role in smoothing the detected anomalies, reducing noise, and providing a more stable and reliable signal. When an anomaly exceeds a configurable threshold, NextSignalSentinel generates a low-latency WebSocket alert, delivering timely information directly to subscribers. This architecture allows for rapid response to market changes, enabling users to capitalize on emerging opportunities and mitigate potential risks. Furthermore, the modular design facilitates easy integration with existing trading infrastructure and custom alert mechanisms.

NextSignalSentinel aims to address the limitations of traditional anomaly detection methods in the crypto market, which often rely on simple moving averages or volume spikes. These methods can be easily fooled by market noise or flash crashes, leading to inaccurate alerts. By combining spectral analysis and Kalman filtering, NextSignalSentinel provides a more nuanced and accurate understanding of market dynamics, enabling users to make more informed trading decisions. The system is designed for high throughput and low latency, ensuring timely delivery of critical information.

## Key Features

*   **Real-time Data Ingestion:** Establishes persistent WebSocket connections to cryptocurrency exchanges (e.g., Binance, Coinbase Pro) to receive real-time price data. Utilizes the `gorilla/websocket` library for efficient handling of WebSocket communication.
*   **Spectral Analysis via FFT:** Employs Fast Fourier Transform (FFT) algorithms, implemented using the `gonum.org/v1/gonum/fft` package, to convert time-series price data into the frequency domain, revealing underlying patterns and anomalies.
*   **Custom Kalman Filter:** Integrates a custom-designed Kalman filter to smooth detected anomalies and reduce noise. The filter's parameters (process noise covariance, measurement noise covariance) are configurable for optimal performance based on the specific cryptocurrency and market conditions.
*   **Configurable Anomaly Thresholds:** Allows users to define thresholds for anomaly detection based on spectral power deviations. Thresholds can be customized for different cryptocurrencies and timeframes.
*   **Low-Latency WebSocket Alerts:** Delivers real-time alerts via WebSocket connections when anomalies are detected. Alerts include detailed information about the anomaly, such as the cryptocurrency pair, timestamp, and anomaly magnitude. Utilizes JSON encoding for efficient data transmission.
*   **Modular Architecture:** Designed with a modular architecture to facilitate easy integration with existing trading infrastructure and custom alert mechanisms. Modules include data ingestion, spectral analysis, Kalman filtering, and alert generation.
*   **Data Persistence (Optional):** Offers optional data persistence using a time-series database (e.g., InfluxDB) for historical analysis and backtesting. Utilizes the InfluxDB Go client for seamless integration.

## Technology Stack

*   **Go:** The primary programming language, chosen for its concurrency capabilities, performance, and suitability for network programming.
*   **gorilla/websocket:** A popular Go library for handling WebSocket connections, providing efficient and reliable communication with cryptocurrency exchanges.
*   **gonum.org/v1/gonum/fft:** A package from the Gonum project, providing optimized implementations of Fast Fourier Transform (FFT) algorithms for spectral analysis.
*   **JSON:** Used for data serialization and deserialization, enabling efficient communication between components and with external systems.
*   **InfluxDB (Optional):** A time-series database used for optional data persistence and historical analysis.

## Installation

1.  **Install Go:** Download and install the Go programming language from the official Go website (https://go.dev/dl/). Ensure that the `GOPATH` and `PATH` environment variables are correctly configured.

2.  **Clone the Repository:** Clone the NextSignalSentinel repository from GitHub:
    `git clone https://github.com/uhsr/NextSignalSentinel.git`

3.  **Navigate to the Project Directory:** Change your current directory to the cloned repository:
    `cd NextSignalSentinel`

4.  **Download Dependencies:** Use the `go mod` command to download the project's dependencies:
    `go mod download`

5.  **Build the Executable:** Build the NextSignalSentinel executable:
    `go build -o nextsignalsentinel`

## Configuration

NextSignalSentinel utilizes environment variables for configuration. Create a `.env` file in the project directory with the following variables:

*   `EXCHANGE_API_URL`: The WebSocket URL of the cryptocurrency exchange (e.g., "wss://stream.binance.com:9443/ws/btcusdt@trade").
*   `ANOMALY_THRESHOLD`: The threshold for anomaly detection (e.g., 0.05). A higher value results in fewer alerts.
*   `KALMAN_PROCESS_NOISE`: The process noise covariance for the Kalman filter (e.g., 0.01).
*   `KALMAN_MEASUREMENT_NOISE`: The measurement noise covariance for the Kalman filter (e.g., 0.1).
*   `ALERT_WEBSOCKET_URL`: The WebSocket URL for sending alerts (optional).
*   `INFLUXDB_URL` (Optional): The URL of the InfluxDB instance for data persistence.
*   `INFLUXDB_TOKEN` (Optional): The authentication token for InfluxDB.
*   `INFLUXDB_ORG` (Optional): The InfluxDB organization.
*   `INFLUXDB_BUCKET` (Optional): The InfluxDB bucket.

Example `.env` file:

EXCHANGE_API_URL=wss://stream.binance.com:9443/ws/btcusdt@trade
ANOMALY_THRESHOLD=0.05
KALMAN_PROCESS_NOISE=0.01
KALMAN_MEASUREMENT_NOISE=0.1

## Usage

1.  **Run the Executable:** Execute the built binary:
    `./nextsignalsentinel`

2.  **Receive Alerts:** If the `ALERT_WEBSOCKET_URL` is configured, NextSignalSentinel will send real-time alerts to the specified WebSocket endpoint. Otherwise, alerts will be printed to the console.

Example alert payload (JSON):

{
"timestamp": "2024-01-01T00:00:00Z",
"symbol": "BTCUSDT",
"anomaly_score": 0.06,
"message": "Anomaly detected"
}

## Contributing

We welcome contributions to NextSignalSentinel! Please follow these guidelines:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Write clear and concise commit messages.
4.  Submit a pull request with a detailed description of your changes.
5.  Ensure that your code adheres to the Go coding style guidelines.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/uhsr/NextSignalSentinel/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to acknowledge the contributions of the open-source community to the Go ecosystem, particularly the developers of the `gorilla/websocket` and `gonum.org/v1/gonum/fft` libraries. Their work has been invaluable in the development of NextSignalSentinel.