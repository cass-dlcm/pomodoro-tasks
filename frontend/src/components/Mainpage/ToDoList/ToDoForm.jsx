import React, { useState, useEffect, useRef } from 'react';
import {gql, useMutation, useQuery} from "@apollo/client";
import styles from "../../SignUp/SignUp.module.scss";

function ToDoForm(props){
    const [input, setInput] = useState('');
    const MUT = gql`mutation CreateTodo($name: String!, $list: ID!) { createTodo(input: {name: $name, list: $list}) {id}}`;
    const [createTodo, { data, loading, error }] = useMutation(MUT);
    const inputRef = useRef(null);

    useEffect(() => {
        inputRef.current.focus();
    });

    const handleChange = e => {
        setInput(e.target.value);
    };

    if (error) {
        return `Error: ${error.message}`;
    }

    if (data) {
        props.list.tasks.push({id: data.createTodo.id, name: input, completedAt: null});
    }

    return (
        <form onSubmit={e => {
            e.preventDefault();
            alert(input);
            createTodo({variables: {name: input, list: props.list.id}})
        }} className='todo-form'>
            <input
                placeholder='Add a To Do'
                value={input}
                onChange={handleChange}
                name='text'
                className='todo-input'
                ref={inputRef}
            />
            <input type={"submit"} value={"Add Task"} className='todo-button'/>
        </form>
    );
}
export default ToDoForm;
