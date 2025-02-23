import rolesAndResources from "../../data/rolesAndResources";
import Profile from "../Profile/Profile";
import styles from "./Sidebar.module.css";
import { Link } from "react-router-dom";  

export default function Sidebar({ userData }) {
  if (!userData) {
    return <div>Loading...</div>;
  }

  const { role, fullname } = userData;
  const Role = role?.toLowerCase() || "";

  const paths = rolesAndResources.get(Role)?.map((path) => (
    <Link key={path} to={path.toLowerCase()}>
      <li>{path}</li>
    </Link>
  )) || [];
  const profile_url = "http://localhost:5000/upload/vishnu-profile.jpg"
  return (
    <div className={styles.sidebar}>
      <Profile image={profile_url} name={fullname} />
      <ul>{paths}</ul>
    </div>
  );
}
