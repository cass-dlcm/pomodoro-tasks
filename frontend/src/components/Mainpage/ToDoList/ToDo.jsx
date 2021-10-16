import React, { useState } from 'react';
import { ToDoCheckbox } from './ToDoCheckbox';

const ToDo = (props) => {
    return <div
            className={`todo-row`}
            key={props.key}
            style={{order: props.index}}
        >
            <ToDoCheckbox todo={props.todo} />
            {props.todo.name}
        </div>
};

export default ToDo;
