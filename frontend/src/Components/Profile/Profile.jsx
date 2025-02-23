import styles from "./Profile.module.css";

export default function Profile({ image, name }) {
  return (
    <div className={styles.profile}>
      <img src={image} alt="Profile" />
      <p>{name}</p>
    </div>
  );
}