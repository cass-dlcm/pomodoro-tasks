import React, { useState, useEffect, useRef } from 'react';
import { gql, useMutation } from "@apollo/client";

function ToDoForm(props){
    const [input, setInput] = useState('');
    const MUT = gql`mutation CreateTodo($name: String!, $list: ID!) { createTodo(input: {name: $name, list: $list}) {id}}`;
    const [createTodo, { data, loading, error }] = useMutation(MUT);
    const inputRef = useRef(null);

    useEffect(() => {
        inputRef.current.focus();
    });

    function handleChange(e) {
        setInput(e.target.value);
    };

    if (error) {
        return `Error: ${error.message}`;
    }

    if (loading) {
        return "Loading...";
    }

    if (data) {
        props.list.tasks.push({id: data.createTodo.id, name: input, completedAt: null});
    }

    function formSubmit(e) {
        e.preventDefault();
        createTodo({variables: {name: input, list: props.list.id}})
    }

    return (
        <form onSubmit={formSubmit} className='todo-form'>
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
