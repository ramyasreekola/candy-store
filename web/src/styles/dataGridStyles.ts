export const dataGridStyles = {
    width: "80%",
    margin: "0 auto",
    border: "none",
    borderRadius: 2,
    boxShadow: "0px 3px 6px rgba(0,0,0,0.29)",
    "& .MuiDataGrid-columnHeaders": {
      backgroundColor: "#a3afc9",
      color: "#fdfefe",
      fontSize: "20px",
      minHeight: "72px !important",
      paddingLeft: "25px",
    },
    "& .MuiDataGrid-cell": {
      borderBottom: "1px solid #d3d3d3",
      fontSize: "16px",
    },
    "& .MuiDataGrid-row": {
      paddingLeft: "25px",
      "&:hover": {
        backgroundColor: "#e0e0e0",
      },
    },
    "& .MuiDataGrid-footerContainer": {
      backgroundColor: "#a3afc9",
      color: "#fdfefe",
      minHeight: "60px !important",
    },
    "& .MuiTablePagination-root": {
      color: "#fdfefe",
    },
  }