import styles from "./SuggestionProfile.module.css"

export default function SuggestionProfile(props){
    return(
        <div className={styles.container}>
            <img src={props.image_url} alt="" className={styles.image}/>
            <p className={styles.fullname}>{props.fullname}</p>
            <button className={styles.button}>Connect</button>
        </div>
    )
}