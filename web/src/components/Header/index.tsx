import React from 'react';
import styles from "./header.module.css";
import CandyIcon from "../../assets/CandyIcon";

const Header: React.FC = () => (
  <div className={styles.header}>
    <CandyIcon />
    <p className={styles.title}>Candy Store</p>
    <CandyIcon />
  </div>
);

export default Header;
