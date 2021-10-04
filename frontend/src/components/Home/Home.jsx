import styles from "./Home.module.scss";
import {useHistory} from "react-router";
import { gql, useMutation } from "@apollo/client";

const MUT = gql`mutation SignIn($name: String!, $pass: String!) { signIn(user: {name: $name, password: $pass})}`

export function Home (props) {
    let history = useHistory();
    let username;
    let password;
    const [signIn, { data, loading, error }] = useMutation(MUT);

    if (loading) return 'Submitting...';
    if (error) return `Submission error! ${error.message}`;

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
                            <h1>Welcome to Stidy</h1>
                        </div>
                        <form onSubmit={e => {
                            e.preventDefault();
                            signIn({variables: {name: username.value, pass: password.value}}).then(r => {
                                props.sendJWT(r.data.toString());
                                history.push('/mainpage')
                            });
                            username.value = '';
                            password.value = '';
                        }}>
                            <div className={styles.inputSection}>
                                <input className={styles.input} ref={node => {
                                    username = node;
                                }} type="text" placeholder="Username"/>
                                <input className={styles.input} ref={node => {
                                    password = node;
                                }} type="text" placeholder="Password"/>
                            </div>
                            <div className={styles.btnSection}>
                                <button type={"submit"} className={styles.btn1}>
                                    Log In
                                </button>
                                <button onClick={() => {history.push("/signup")}} type="button" className={styles.btn2}>
                                    Sign Up
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    );
}