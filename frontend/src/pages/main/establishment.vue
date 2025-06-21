<template>
    <div class="w-full flex flex-col gap-4">
        <div class="w-full flex gap-2 items-center justify-center">
            <h1 class="font-bold">ESTABELECIMENTOS</h1>
            <NuxtButton variant="outline" @click="openCreating">New</NuxtButton>
        </div>

        <div class="flex flex-col">
            <NuxtTable :data="data" :columns="columns" class="flex-1" />
            <div class="w-full flex px-6 justify-end">
                <NuxtPagination v-model:page="page" :total="100" class="end" /> 
            </div>
        </div>
    </div>
    
    <NuxtModal v-model:open="creating">
        <template #content>
            <Placeholder class="m-4 p-4">
                <NuxtForm :schema="schema" :state="state" class="space-y-4" @submit="newEstablishment">
                    <div class="flex flex-col rounded-md gap-8 items-center">
                        <div class="w-full flex gap-4">
                            <NuxtFormField label="Nome" name="name" required>
                                <NuxtInput v-model="state.name" />
                            </NuxtFormField>
                            
                            <NuxtFormField label="Razão Social" name="legal_name" required>
                                <NuxtInput v-model="state.legal_name" />
                            </NuxtFormField>
    
                            <NuxtFormField label="Numero" name="number" required>
                                <NuxtInput v-model="state.number" />
                            </NuxtFormField>
                        </div>

                        <div class="w-full grid grid-cols-3 gap-4 justify-between">
                            <NuxtFormField label="Rua" name="street">
                                <NuxtInput v-model="state.street" />
                            </NuxtFormField>

                            <NuxtFormField label="Numero da Rua" name="address_number">
                                <NuxtInput v-model="state.address_number" />
                            </NuxtFormField>

                            <NuxtFormField label="Cidade" name="city" required>
                                <NuxtInput v-model="state.city" />
                            </NuxtFormField>
    
                            <NuxtFormField label="Estado" name="state" required>
                                <NuxtInput v-model="state.state" />
                            </NuxtFormField>
    
                            <NuxtFormField label="CEP" name="zip_code">
                                <NuxtInput v-model="state.zip_code" />
                            </NuxtFormField>
                        </div>

                        <div class="w-full flex items-center justify-evenly">
                            <NuxtButton type="submit">New</NuxtButton>
                            <NuxtButton variant="outline" @click="closeCreating">Cancel</NuxtButton>
                        </div>
                    </div>
                </NuxtForm>
            </Placeholder>
        </template>
    </NuxtModal>

    <NuxtModal v-model:open="isModalOpen">
        <template #content>
            <Placeholder class="m-4 p-4">
                <NuxtForm :schema="schema" :state="state" class="space-y-4" @submit="updateEstablishment">
                    <div class="flex flex-col rounded-md gap-8 items-center">
                        <div class="w-full flex gap-4">
                            <NuxtFormField label="Nome" name="name" required>
                                <NuxtInput v-model="state.name" />
                            </NuxtFormField>
                            
                            <NuxtFormField label="Razão Social" name="legal_name" required>
                                <NuxtInput v-model="state.legal_name" />
                            </NuxtFormField>
    
                            <NuxtFormField label="Numero" name="number" required>
                                <NuxtInput v-model="state.number" />
                            </NuxtFormField>
                        </div>

                        <div class="w-full grid grid-cols-3 gap-4 justify-between">
                            <NuxtFormField label="Rua" name="street">
                                <NuxtInput v-model="state.street" />
                            </NuxtFormField>

                            <NuxtFormField label="Numero da Rua" name="address_number">
                                <NuxtInput v-model="state.address_number" />
                            </NuxtFormField>

                            <NuxtFormField label="Cidade" name="city" required>
                                <NuxtInput v-model="state.city" />
                            </NuxtFormField>
    
                            <NuxtFormField label="Estado" name="state" required>
                                <NuxtInput v-model="state.state" />
                            </NuxtFormField>
    
                            <NuxtFormField label="CEP" name="zip_code">
                                <NuxtInput v-model="state.zip_code" />
                            </NuxtFormField>
                        </div>

                        <div class="w-full flex items-center justify-evenly">
                            <NuxtButton type="submit">Save</NuxtButton>
                            <NuxtButton variant="outline" @click="closeModal">Cancel</NuxtButton>
                        </div>
                    </div>
                </NuxtForm>
            </Placeholder>
        </template>
    </NuxtModal>
</template>

<script setup lang="ts">
    definePageMeta({
        // layout: '/',
    });

    import * as z from 'zod';
    import { ref, h, resolveComponent } from 'vue';
    import type { TableColumn } from '@nuxt/ui'
    import type { Row } from '@tanstack/vue-table';
    import { useClipboard } from '@vueuse/core';

    import { useEstablishment } from '~/composables/useEstablishment';
    import type { EstablishmentPreview } from "~/types/establishment";
    
    const NuxtButton = resolveComponent('NuxtButton');
    const NuxtDropdownMenu = resolveComponent('NuxtDropdownMenu');
    
    const toast = useToast();
    const { copy } = useClipboard();
    
    
    const { create, list, update, remove } = useEstablishment();
    
    const { data: establishments, error } = await useAsyncData(
        'establishments',
        async () => await list(),
    );
    
    if (error.value) {
        console.error('Erro ao carregar estabelecimentos:', error.value);
    }
    
    const page = ref(1);

    const isModalOpen = ref<boolean>(false);
    const creating = ref<boolean>(false);
    const establishmentToUpdate = ref<EstablishmentPreview | null>(null);

    const schema = z.object({
        id: z.string(),
        name: z.string().min(3, "Campo obrigatório"),
        legal_name: z.string().min(3, "Campo obrigatório"),
        number: z.string().min(1, "Campo obrigatório"),
        
        street: z.string(),
        address_number: z.string(),
        city: z.string().min(2, "Campo obrigatório"),
        state: z.string().min(2, "Campo obrigatório"),
        zip_code: z.string(),
    });

    type Schema = z.output<typeof schema>;

    const state = reactive<Partial<Schema>>({
        id:'',
        name:'',
        legal_name:'',
        number:'',
        street:'',
        address_number:'',
        city:'',
        state:'',
        zip_code:'',
    });

    watch(establishmentToUpdate, (newEstablishment: EstablishmentPreview) => {
        if (newEstablishment) {
            state.id = newEstablishment.id;
            state.name = newEstablishment.name;
            state.legal_name = newEstablishment.legal_name;
            state.number = newEstablishment.number;
            state.street = newEstablishment.street;
            state.address_number = newEstablishment.address_number;
            state.city = newEstablishment.city;
            state.state = newEstablishment.state;
            state.zip_code = newEstablishment.zip_code;
        }
    });

    const openCreating = async (establishment: EstablishmentPreview) => {
        Object.assign(state, {
            id: '',
            name: '',
            legal_name: '',
            number: '',
            street: '',
            address_number: '',
            city: '',
            state: '',
            zip_code: '',
        });

        establishmentToUpdate.value = null;
        creating.value = true;
    }

    const closeCreating = async () => {
        creating.value = false;
    }

    const openModal = async (establishment: EstablishmentPreview) => {
        establishmentToUpdate.value = establishment;
        isModalOpen.value = true;
    }

    const closeModal = async () => {
        isModalOpen.value = false;
    }

    function mapEstablishmentData(establishments: EstablishmentPreview[]) {
        return establishments.map((establishment: EstablishmentPreview): DataPreview => ({
            id: establishment.id,
            name: establishment.name,
            legal_name: establishment.legal_name,
            number: establishment.number,
            street: establishment.address.street,
            address_number: establishment.address.number,
            city: establishment.address.city,
            state: establishment.address.state,
            zip_code: establishment.address.zip_code,
        }));
    }


    const newEstablishment = async (form: { data: Schema }) => {
        const values = form.data

        const payload = {
            name: values.name,
            legal_name: values.legal_name,
            number: values.number,
            address: {
                street: values.street,
                number: values.address_number,
                city: values.city,
                state: values.state,
                zip_code: values.zip_code,
            }
        }

        try {
            await create(payload)

            const freshList = await list()
            data.value = mapEstablishmentData(freshList)

            toast.add({ title: 'Estabelecimento criado com sucesso!', color: 'success' })
            creating.value = false
        } catch (err) {
            toast.add({ title: 'Erro ao criar estabelecimento', color: 'red' })
            console.error(err)
        }
    }

    const updateEstablishment = async (establishment: Partial<Schema>) => {
        const values = establishment.data;
        try {
            const updatePayload = {
                name: values.name,
                legal_name: values.legal_name,
                number: values.number,
                address: {
                    street: values.street,
                    number: values.address_number,
                    city: values.city,
                    state: values.state,
                    zip_code: values.zip_code,
                }
            };

            if (!values.id) {
                throw new Error('ID do estabelecimento não informado');
            }

            await update(values.id, updatePayload);

            const freshList = await list();
            data.value = mapEstablishmentData(freshList);

            toast.add({ title: 'Estabelecimento atualizado com sucesso!', color: 'success' });
            isModalOpen.value = false;
        } catch (err) {
            toast.add({ title: 'Erro ao atualizar estabelecimento', color: 'red' });
            console.error(err);
        }
    }


    const removeAndRefresh = async (id: string) => {
        try {
            await remove(id);
            const freshList = await list();
            data.value = mapEstablishmentData(freshList);
            toast.add({ title: 'Establishment deleted.', color: 'success' });
        } catch (err) {
            toast.add({ title: 'Erro ao deletar', color: 'red' });
            console.error(err);
        }
    };

    type DataPreview = {
        id: string;
        name: string
        legal_name: string
        number: string
        street: string
        address_number: string
        city: string
        state: string
        zip_code: string
    }

    const establishmentsPreview = mapEstablishmentData(establishments.value);

    // CONFIGURAÇÃO DAS COLUMNS
    const columns: TableColumn<DataPreview>[] = [
        {
            accessorKey: 'name',
            header: () => h('fiv', { class: 'text-right' }, 'Nome'),
            cell: ({ row }) => h('div', { class: 'text-left font-medium' }, row.getValue('name')),
        },
        {
            accessorKey: 'legal_name',
            header: () => h('fiv', { class: 'text-right' }, 'Razão Social'),
            cell: ({ row }) => h('div', { class: 'text-left font-medium' }, row.getValue('legal_name')),
        },
        {
            accessorKey: 'number',
            header: () => h('fiv', { class: 'text-right' }, 'Numero'),
            cell: ({ row }) => h('div', { class: 'text-left font-medium' }, row.getValue('number')),
        },
        {
            accessorKey: 'street',
            header: () => h('fiv', { class: 'text-right' }, 'Rua'),
            cell: ({ row }) => h('div', { class: 'text-left font-medium' }, row.getValue('street'))
        },
        {
            accessorKey: 'address_number',
            header: () => h('fiv', { class: 'text-right' }, 'Numero Endereço'),
            cell: ({ row }) => h('div', { class: 'text-left font-medium' }, row.getValue('address_number'))
        },
        {
            accessorKey: 'city',
            header: () => h('fiv', { class: 'text-right' }, 'Cidade'),
            cell: ({ row }) => h('div', { class: 'text-left font-medium' }, row.getValue('city'))
        },
        {
            accessorKey: 'state',
            header: () => h('fiv', { class: 'text-right' }, 'Estado'),
            cell: ({ row }) => h('div', { class: 'text-left font-medium' }, row.getValue( 'state'))
        },
        {
            accessorKey: 'zip_code',
            header: () => h('fiv', { class: 'text-right' }, 'CEP'),
            cell: ({ row }) => h('div', { class: 'text-left font-medium' }, row.getValue('zip_code'))
        },
        {
            id: 'actions',
            cell: ({ row }) => {
            return h(
                'div',
                { class: 'text-right text-white' },
                h(
                    NuxtDropdownMenu,
                    {
                        items: getRowItems(row),
                        'aria-label': 'Actions dropdown',
                    },
                    () =>
                        h(NuxtButton, {
                            icon: 'i-lucide-ellipsis-vertical',
                            color: 'neutral',
                            variant: 'ghost',
                            class: 'ml-auto',
                            'aria-label': 'Actions dropdown',
                        })
                    )
                )
            }
        }
    ];
    
    function getRowItems(row: Row<DataPreview>) {
        return [
            {
                type: 'label',
                label: 'Actions'
            },
            {
                label: 'Copy establishment ID',
                onSelect() {
                    copy(row.original.id);
                    toast.add({
                        title: 'Establishment ID copied to clipboard!',
                        color: 'success',
                        icon: 'i-lucide-circle-check',
                    });
                }
            },
            {
                type: 'separator',
            },
            {
                label: 'View establishment details',
                onSelect() {
                    openModal(row.original);
                }
            },
            {
                label: 'Delete establishment',
                onSelect() {
                    removeAndRefresh(row.original.id);
                }
            }
        ];
    }

    const data = ref<DataPreview[]>(establishmentsPreview);
</script>
