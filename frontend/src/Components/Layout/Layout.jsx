import { useEffect, useState } from "react";
import styles from "./Layout.module.css";
import Sidebar from "../Sidebar/Sidebar";
import { Outlet } from "react-router-dom";
import { getUserInfo } from "../../api/requests";

export default function Layout() {
  const [userData, setUserData] = useState(null);
  useEffect(() => {
    setUserData(getUserInfo())
  }, []);

  return (
    <div className={styles.layout}>
      <Sidebar userData={userData} />
      <main>
        <Outlet />
      </main>
    </div>
  );
}