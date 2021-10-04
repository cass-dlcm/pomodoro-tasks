import {gql, useMutation} from "@apollo/client";
import React from "react";

export const ToDoCheckbox = (props) => {
    const MUT = gql`mutation CompleteTodo($id: ID!) {markCompletedTodo(id: $id){completedAt}}`;
    const [completeTodo, { data, loading, error }] = useMutation(MUT, {variables: {id: props.todo.id}});
    if (loading) return 'Loading...';
    if (error) return `Error! ${error.message}`;
    if (data) {
        props.todo.completedAt = data
    }
    return (
        <button onClick={() => completeTodo()} value={"Complete"} >
                <span className="material-icons-outlined">
                    check_box_outline_blank
                </span>
        </button>
    )
};