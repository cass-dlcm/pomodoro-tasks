import styles from "./SignUp.module.scss";
import {Redirect, useHistory} from "react-router";
import { gql, useMutation } from "@apollo/client";

const MUT = gql`mutation CreateUser($name: String!, $password: String!) { createUser(user: {name: $name, password: $password}){id}}`;

export function SignUp () {
    let username;
    let password;
    const [signUp, { data, loading, error }] = useMutation(MUT);

    if (loading) return 'Submitting...';
    if (error) return `Submission error! ${error.message}`;
    if (data) return <Redirect to={'/'} />;

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
                        <form onSubmit={e => {
                            e.preventDefault();
                            signUp({ variables: { name: username.value, password: password.value } });
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
                                <input type={"submit"} value={"Submit"} className={styles.btn1}/>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    );
}
