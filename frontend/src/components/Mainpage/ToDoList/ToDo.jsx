import React, { useState } from 'react';
import { ToDoCheckbox } from './ToDoCheckbox';

const ToDo = (props) => {
    return <div
            className={props.todo.completedAt ? 'todo-row complete' : 'todo-row'}
            key={props.key}
        >
            {props.todo.text}
            <ToDoCheckbox todo={props.todo} />
        </div>
};

export default ToDo;
