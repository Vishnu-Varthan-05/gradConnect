import rolesAndResources from "../../data/rolesAndResources";
import Profile from "../Profile/Profile";
import styles from "./Sidebar.module.css";
import { Link } from "react-router-dom";

export default function Sidebar({ userData }) {
  if (!userData) {
    return <div>Loading...</div>;
  }

  const { profileUrl, userType, username, userId } = userData;

  const paths = rolesAndResources.get(userType).map((path) => (
    <Link to={path.toLowerCase()}>
        <li key={path}>{path}</li>
    </Link>
  ));

  return (
    <div className={styles.sidebar}>
      <Profile image={profileUrl} name={username} />
      <ul>{paths}</ul>
    </div>
  );
}