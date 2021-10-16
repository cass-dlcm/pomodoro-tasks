import React, {useState} from 'react';
import ToDoForm from './ToDoForm';
import ToDo from './ToDo';
import {useQuery, gql} from "@apollo/client";

export const ToDoList = (props) => {
    const [skip, setSkip] = useState(false),
        QUERYTWO = gql`query GetTodos($listId: ID!) {todos(list: $listId){name, tasks {id, name, completedAt}}}`, {
            data: data,
            loading: loading,
            error: error
        } = useQuery(QUERYTWO, {variables: {listId: props.list.id}, skip: skip});
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
            <div className={'todo-container'}>
                {data.todos.tasks.map((todo, index) => (
                    <ToDo
                        key={todo.id}
                        index={index}
                        todo={todo}
                    />
                ))}
            </div>
        </>
    }

    if (props.list.tasks) {
        return <>
            <ToDoForm list={props.list}/>
            <div className={'todo-container'}>
                {props.list.tasks.map((todo, index) => (
                    <ToDo
                        key={todo.id}
                        index={index}
                        todo={todo}
                    />
                ))}
            </div>
        </>
    }

    return <ToDoForm list={props.list}/>;
}