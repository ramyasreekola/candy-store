import React, { useState, useEffect } from "react";
import { Purchase } from "../../types/type";
import styles from "./home.module.css";
import { DataGrid } from "@mui/x-data-grid";
import Header from "../../components/Header";
import { Error as ErrorComponent } from "../404";
import { columns, fetchCandyData } from "../../utils";
import { dataGridStyles } from "../../styles/dataGridStyles";

const Home: React.FC = () => {
  const [purchase, setPurchase] = useState<Purchase[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchCandyData(setPurchase, setError);
  }, []);


  if (error != null) {
    return <ErrorComponent message={error} />;
  }

  return (
    <div className={styles.container}>
      <Header />
      <DataGrid
        rows={purchase}
        columns={columns}
        rowHeight={60}
        sx={dataGridStyles}
        initialState={{
          pagination: {
            paginationModel: { page: 0, pageSize: 10 },
          },
        }}
        pageSizeOptions={[5, 10, 15]}
      />
    </div>
  );
};

export default Home;
