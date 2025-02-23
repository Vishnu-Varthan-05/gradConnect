import Post from "../Post/Post";
import styles from "./Home.module.css";
import SuggestionProfile from "../SuggestionProfile/SuggestionProfile";

export default function Home() {
    const listPosts = [
        {
            id: 1,
            imageUrl: "src/assets/r27vhdprnuv21.jpg",
            title: "First Post",
            description: "This is the description for the first post."
        },
        {
            id: 2,
            imageUrl: "src/assets/rdoedtlvu0x61.png",
            title: "Second Post",
            description: "This is the description for the second post."
        },
        {
            id: 3,
            imageUrl: "src/assets/wallhaven-9dvl38.png",
            title: "Third Post",
            description: "This is the description for the third post."
        }
    ];
    
    const suggestions = [
        {
            id: 1,
            image_url: "src/assets/r27vhdprnuv21.jpg",
            fullname: "John Doe"
        },
        {
            id: 2,
            image_url: "src/assets/rdoedtlvu0x61.png",
            fullname: "Jane Smith"
        },
        {
            id: 3,
            image_url: "src/assets/wallhaven-9dvl38.png",
            fullname: "Alice Johnson"
        }

    ];
    return (
        <div className={styles.homeContainer}>
            <div className={styles.post}>
                {listPosts.map((post) => (
                    <Post 
                        key={post.id} 
                        imageUrl={post.imageUrl} 
                        title={post.title} 
                        description={post.description} 
                    />
                ))}
            </div>
            <div className={styles.suggestion}>
                {suggestions.map((profile) => (
                    <SuggestionProfile 
                        key={profile.id} 
                        image_url={profile.image_url} 
                        fullname={profile.fullname} 
                    />
                ))}
            </div>
        </div>
    );
}
