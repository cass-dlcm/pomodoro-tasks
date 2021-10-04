import React, {useState} from 'react';
import ToDoForm from './ToDoForm';
import ToDo from './ToDo';
import {useMutation, useQuery, gql} from "@apollo/client";
import {ToDoCheckbox} from "./ToDoCheckbox";

export const ToDoList = (props) => {
    const [skip, setSkip] = useState(false);
    const QUERYTWO = gql`query GetTodos($listId: ID!) {todos(list: $listId){name, tasks {id, name, completedAt}}}`
    const {data, loading, error} = useQuery(QUERYTWO, {variables: {listId: props.list.id}, skip: skip})
    if (loading) return 'Loading list items...';
    if (error) return `Error! ${error.message}`;
    if (data) {
        props.setList({id: props.list.id, name: data.todos.name, tasks: data.todos.tasks});
        setSkip(true);
    }
        // const removeTodo = id => {
        //   const MUT = gql`mutation RemoveTodo($id: ID!) {deleteTodo(id: $id)}`
        //   const { data, loading, error } = useMutation(MUT, {variables: {id: id}});
        //   if (loading) return 'Loading...';
        //   if (error) return `Error! ${error.message}`;
        //   if (data) {
        //     props.list.tasks = [...props.list.tasks].filter(todo => todo.id !== id);
        //   }
        // };

    if (data) {
        return <>
            <ToDoForm list={props.list}/>
            {data.todos.tasks.map((todo) => (
                <ToDo
                    key={todo.id}
                    todo={todo}
                />
            ))}
        </>
    }

    if (props.list.tasks) {
        return <>
            <ToDoForm list={props.list}/>
            {props.list.tasks.map((todo) => (
                <ToDo
                    key={todo.id}
                    todo={todo}
                />
            ))}
        </>
    }

    return (
        <>
            <ToDoForm list={props.list}/>
        </>
    );
}