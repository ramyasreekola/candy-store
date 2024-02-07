
import { GridColDef, GridValueGetterParams } from "@mui/x-data-grid";
import { formatDate } from "./format";

export const columns: GridColDef[] = [
    { field: "name", headerName: "Name", width: 260 },
    { field: "candy", headerName: "Candy", width: 260 },
    {
      field: "eaten",
      headerName: "Eaten",
      type: "number",
      width: 200,
      headerAlign: "left",
      align: "left",
    },
    {
      field: "date",
      headerName: "Date",
  
      width: 260,
      valueGetter: (params: GridValueGetterParams) => formatDate(params.row.date),
    },
  ];
  