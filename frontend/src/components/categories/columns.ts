import { h } from 'vue'
import { type ColumnDef } from '@tanstack/vue-table';
import type { Category } from '~/types/category'

export const columns: ColumnDef<Category>[] = [
    {
        accessorKey: 'Name',
        header: () => h('div', { class: 'text-right' }, 'Name'),
        cell: ({ row }) => {
            const name = row.getValue('Name');
            return h('div', { class: 'text-right font-medium' }, String(name));
        },
    },
    {
        accessorKey: 'CreatedAt',
        header: () => h('div', { class: 'text-right' }, 'Created'),
        cell: ({ row }) => {
            const createdAt = String(row.getValue('CreatedAt'));
            const formattedDate = new Date(createdAt).toDateString();
            return h('div', { class: 'text-right font-medium' }, formattedDate);
        },
    },
    {
        accessorKey: 'UpdatedAt',
        header: () => h('div', { class: 'text-right' }, 'Last Update'),
        cell: ({ row }) => {
            const updatedAt = String(row.getValue('UpdatedAt'));
            const formattedDate = new Date(updatedAt).toDateString();
            return h('div', { class: 'text-right font-medium' }, formattedDate);
        },
    },
    // {
    //     id: 'actions',
    //     enableHiding: false,
    //     cell: ({ row }) => {
    //         const category = row.original;
    //         return h('div', { class: 'relative' }, h(DropdownAction, {
    //             category,
    //         }))
    //     },
    // },
    // {
    //     accessorKey: 'price',
    //     header: () => h('div', { class: 'text-right' }, 'Price'),
    //     cell: ({ row }) => {
    //         const amount = (Number.parseFloat(row.getValue('price')) / 100.0);
    //         const formatted = new Intl.NumberFormat('en-US', {
    //             style: 'currency',
    //             currency: 'USD',
    //         }).format(amount);

    //         return h('div', { class: 'text-right font-medium' }, formatted)
    //     },
    // }
];
