import React, { useState } from 'react';
import { ToDoCheckbox } from './ToDoCheckbox';

const ToDo = (props) => {
    return <div
            className={props.todo.completedAt ? 'todo-row complete' : 'todo-row'}
            key={props.key}
        >
            <ToDoCheckbox todo={props.todo} />
            {props.todo.text}
        </div>
};

export default ToDo;
