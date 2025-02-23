import styles from "./Post.module.css"

export default function Post(props) {    
    return(
        <div className={styles.postContainer}>
            <img src={props.imageUrl} alt=""  className={styles.postImage}/>
            <p className={styles.postTitle}>{props.title}</p>
            <p className={styles.postDescription}>{props.description}</p>
        </div>
    );
}