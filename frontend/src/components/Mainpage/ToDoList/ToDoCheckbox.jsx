import {gql, useMutation} from "@apollo/client";
import React from "react";

export const ToDoCheckbox = (props) => {
    const MUT = gql`mutation CompleteTodo($id: ID!) {markCompletedTodo(id: $id){completedAt}}`;
    const [completeTodo, { data, loading, error }] = useMutation(MUT, );
    if (loading) return 'Loading...';
    if (error) return `Error! ${error.message}`;
    if (data) {
        props.todo.completedAt = data
    }
    return (
        <button onClick={() => completeTodo({variables: {id: props.todo.id}, skip: props.todo.completedAt})} value={"Complete"} >
                <span className="material-icons">
                    {props.todo.completedAt ? 'check_box' : 'check_box_outline_blank'}
                </span>
        </button>
    )
};