import {gql, useMutation} from "@apollo/client";
import React from "react";

export const ToDoCheckbox = (props) => {
    const MUT = gql`mutation CompleteTodo($id: ID!) {markCompletedTodo(id: $id){completedAt}}`, [completeTodo, {
        data: data,
        loading: loading,
        error: error
    }] = useMutation(MUT,);
    if (loading) return 'Loading...';
    if (error) return `Error! ${error.message}`;
    if (data) {
        props.todo.completedAt = data
    }

    function checkboxClicked() {
        if (props.todo.completedAt == null) completeTodo({variables: {id: props.todo.id}, skip: props.todo.completedAt})
    }

    return (
        <button onClick={checkboxClicked} value={"Complete"} >
                <span className="material-icons">
                    {props.todo.completedAt ? 'check_box' : 'check_box_outline_blank'}
                </span>
        </button>
    )
};