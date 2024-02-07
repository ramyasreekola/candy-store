# Candy Store Dashboard
This project is a React-based frontend application using Go backend showcasing candy consumption data in a user-friendly dashboard.

## Backend 

The backend service provides access to candy consumption data via a RESTful API. It is built using the Echo framework in Go, offering a single endpoint that returns a collection of candy consumption records.

### Built With
* [Go](https://go.dev/) - The Go Programming Language.
* [Echo](https://echo.labstack.com/) - High performance, extensible, minimalist Go web framework for backend API data fetching.
* [Zimpler Candy Store API](https://candystore.zimpler.net/api/candy) - Candy data fetch.

### Features
* Fetch candy consumption data from an [external API](https://candystore.zimpler.net/api/candy).
* Serve fetched data through a RESTful API.
* Error handling for the data fetch fail from external API and incase of time data from api is not formattable.
* Custom time parsing to handle date formats.
* Unit tests added for the custom timeFormat and `fetchCandyData` functions

#### Prerequisites
Before you begin, ensure you have installed Go (version 1.16 or later is recommended) on your system. You can download it from https://golang.org/dl/.

#### Running the Service
To start the service, clone the code and run:

```bash=
git clone https://github.com/ramyasreekola/candy-store
go run main.go
```
This will start the server on http://localhost:3000. The service can be accessed via the `/api/candy` endpoint.


## Frontend
Built with Vite, the app efficiently loads and displays candy data fetched from an external API, presenting it in an interactive and accessible format. The frontend application code  and its configurations, dependencies are located within the `web` sub-folder from the main directory.


### Built With
[React](https://react.dev/) - A JavaScript library for building user interfaces.
[Vite](https://vitejs.dev/) - Next Generation Frontend Tooling.
[MUI](https://mui.com/) - Google's react based open source UI framework.
[DataGrid](https://mui.com/x/react-data-grid/) - Material UI's Data Grid react component



### Getting Started
These instructions will get you a copy of the project up and running locally.

#### Installation
Clone the repository:
```bash
cd candy-store/web
```
Install dependencies and start dev server:

```bash
npm install
npm run dev
```

The application will be available at http://localhost:5713. (vite's default application port)

### Components & Views
In this Application we have the following page views and components
**Home**: Main dashboard view displaying the candy consumption data.
**Error**: Error display component for handling fetch-related errors.
**NotFound**: Page Not Found component to gracefully visualize the routes not supported.
**Header**: To display a short header description of the dashboard.

### API
This frontend app fetches data from the `/api/candy` endpoint, which is expected to return a JSON array of candy consumption records.This api is request is served from the backend service setup and running earlier.

### Features
* Display a list of candy consumption records in a clean and precise tabular format.
* Responsive design for optimal viewing on various devices and screen sizes.
* Error handling to gracefully manage and notify users of any data fetching issues.
* Uses MUI's DataGrid which is a robust, feature rich tabular data component.
* Supports pagination and custom setting of number of rows per page.
* Possible to filter on all the table columns and show relevant filtered data.
* Sorting supported for all column headers(ascending or descending for numeric cell values and chronologival order for cells containing string values).

#### TODO

* For a larger application,styles that will be reused across the application, a custom MUI theme can be added to style the Data Grid table we are using. 
* Add unit tests for the frontend code.
* Suspence loading or a placeholder UI while the api data is loading.


Authors

[Ramya Sree Kola](https://ramyasreekola.github.io/Portfolio/)



