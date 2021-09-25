import styles from "./Home.module.scss";

export const Home = () => {
    return (
        <div>
            <div className={styles.background}>
                <div className={styles.sec1}>
                <img src="./assets/Laptop.png" alt="" />
                </div>
                <div className={styles.sec2}>
                <img src="./assets/StickyNote.png" alt="" />
                <div className={styles.loginSection}>
                <div className="heading">
                <h1>Welcome to Stidy</h1>
                </div>
                <div>
                    <div className="input-section">
                    <input type="text" />
                    <input type="text" />
                    </div>
                    <div className="btn-section">
                        <button>Log In</button>
                        <button>Sign Up</button>
                    </div>
                </div>
                

                </div>
                </div>
            </div>
        </div>
    );
}