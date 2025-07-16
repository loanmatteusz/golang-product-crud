import { h } from 'vue'
import { type ColumnDef } from '@tanstack/vue-table';
import type { Category } from '~/types/category'
import Checkbox from '../ui/checkbox/Checkbox.vue';
import DataTableDropdown from './data-table-dropdown.vue';

export const columns: ColumnDef<Category>[] = [
    {
        id: 'select',
        header: ({ table }) => h(Checkbox, {
            'modelValue': table.getIsAllPageRowsSelected(),
            'onUpdate:modelValue': (value: boolean) => table.toggleAllPageRowsSelected(!!value),
            'ariaLabel': 'Select all',
        }),
        cell: ({ row }) => h(Checkbox, {
            'modelValue': row.getIsSelected(),
            'onUpdate:modelValue': (value: boolean) => row.toggleSelected(!!value),
            'ariaLabel': 'Select row',
        }),
        enableSorting: false,
        enableHiding: false,
    },
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
    {
        id: 'actions',
        enableHiding: false,
        header: () => h('div', { class: 'text-right' }, 'Actions'),
        cell: ({ row }) => {
            const category = row.original;
            return h('div', { class: 'relative text-right' }, h(DataTableDropdown, {
                category,
            }))
        },
    },
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
