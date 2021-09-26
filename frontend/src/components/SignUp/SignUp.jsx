import styles from "./SignUp.module.scss";

export const SignUp = () => {
    return (
        <div>
            <div className={styles.background}>
                <div className={styles.sec1}>
                <img src="./assets/Laptop.png" alt="" />
                </div>
                <div className={styles.sec2}>
                <img src="./assets/StickyNote.png" alt="" />
                <div className={styles.loginSection}>
                <div className={styles.heading}>
                <h1>Sign Up</h1>
                </div>
                <div>
                    <div className={styles.inputSection}>
                    <input className={styles.input} type="text" placeholder="Username"/>
                    <input className={styles.input} type="text" placeholder="Password"/>
                    </div>
                    <div className={styles.btnSection}>
                        <button className={styles.btn1}>Register!</button>
                    </div>
                </div>
                </div>
                </div>
            </div>
        </div>
    );
}