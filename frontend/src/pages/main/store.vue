<template>
    <div class="w-full flex flex-col text-center">
        <div class="w-full flex gap-2 items-center justify-center">
            <h1 class="font-bold">LOJAS</h1>
            <NuxtButton variant="outline" @click="openCreating">New</NuxtButton>
        </div>
        <div class="flex flex-col">
            <NuxtTable :data="data" :columns="columns" class="flex-1" />
            <div class="w-full flex px-6 justify-end">
                <NuxtPagination v-model:page="page" :total="100" class="end" /> 
            </div>
        </div>
    </div>


    <!-- // CRIAÇÃO -->
    <NuxtModal v-model:open="creating">
        <template #content>
            <Placeholder class="m-4 p-4">
                <NuxtForm :schema="schema" :state="state" class="space-y-4" @submit="createStore">
                    <div class="flex flex-col rounded-md gap-8 items-center">
                        <div class="w-full flex">
                            <NuxtFormField label="Estabelecimento">
                                <NuxtSelect v-model="state.establishment_id" :items="selectItems" class="w-48" placeholder="Choose an point" />
                            </NuxtFormField>
                        </div>

                        <div class="w-full flex gap-4">
                            <NuxtFormField label="Nome" name="name" required>
                                <NuxtInput v-model="state.name" placeholder="nome" />
                            </NuxtFormField>
                            
                            <NuxtFormField label="Razão Social" name="legal_name" required>
                                <NuxtInput v-model="state.legal_name" placeholder="razao social" />
                            </NuxtFormField>

                            <NuxtFormField label="Numero" name="number" required>
                                <NuxtInput v-model="state.number" placeholder="numero" />
                            </NuxtFormField>
                        </div>

                        <div class="w-full grid grid-cols-3 gap-4 justify-between">
                            <NuxtFormField label="Rua" name="street">
                                <NuxtInput v-model="state.street" placeholder="nome da rua" />
                            </NuxtFormField>

                            <NuxtFormField label="Numero da Rua" name="address_number">
                                <NuxtInput v-model="state.address_number" placeholder="numero do endereço" />
                            </NuxtFormField>

                            <NuxtFormField label="Cidade" name="city" required>
                                <NuxtInput v-model="state.city" placeholder="cidade" />
                            </NuxtFormField>

                            <NuxtFormField label="Estado" name="state" required>
                                <NuxtInput v-model="state.state" placeholder="estado" />
                            </NuxtFormField>

                            <NuxtFormField label="CEP" name="zip_code">
                                <NuxtInput v-model="state.zip_code" placeholder="cep" />
                            </NuxtFormField>
                        </div>

                        <div class="w-full flex items-center justify-evenly">
                            <NuxtButton type="submit">Criar</NuxtButton>
                            <NuxtButton variant="outline" @click="closeCreating">Cancel</NuxtButton>
                        </div>
                    </div>
                </NuxtForm>
            </Placeholder>
        </template>
    </NuxtModal>


    
    <!-- update -->
    <NuxtModal v-model:open="isModalOpen">
        <template #content>
            <Placeholder class="m-4 p-4">
                <NuxtForm :schema="schema" :state="state" class="space-y-4" @submit="updateStore">
                    <div class="flex flex-col rounded-md gap-8 items-center">
                        <div class="w-full flex">
                            <NuxtFormField label="Estabelecimento">
                                <NuxtSelect v-model="state.establishment_id" :items="selectItems" class="w-48" />
                            </NuxtFormField>
                        </div>

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
    import * as z from 'zod/v4';
    import { ref, h, resolveComponent } from 'vue';
    import type { TableColumn } from '@nuxt/ui'
    import type { Row } from '@tanstack/vue-table';
    import { useClipboard } from '@vueuse/core';

    import { useStore } from '~/composables/useStore';
    import type { StorePreview } from "~/types/store";
    import type { EstablishmentPreview } from '~/types/establishment';
    
    const NuxtButton = resolveComponent('NuxtButton');
    const NuxtDropdownMenu = resolveComponent('NuxtDropdownMenu');
    
    const toast = useToast();
    const { copy } = useClipboard();
    
    
    const { create, list, update, remove } = useStore();
    const { list: listEstablishments } = useEstablishment();
    
    const { data: stores, error: storeError } = await useAsyncData(
        'stores',
        async () => await list(),
    );

    const { data: establishments, error: establishmentError } = await useAsyncData(
        'establishments',
        async () => await listEstablishments(),
    );
    
    if (storeError.value || establishmentError.value) {
        console.error('Erro ao carregar dados:', storeError.value);
    }
    
    const page = ref(1);
    const creating = ref(false);
    
    const isModalOpen = ref<boolean>(false);
    const storeToUpdate = ref<Partial<StorePreview> | null>(null);
    
    const mapSelectItems = (estabs: EstablishmentPreview[]) => (estabs.map(estab => ({
        label: estab.name,
        value: estab.id,
    })));

    const selectItems = ref(mapSelectItems(establishments.value as EstablishmentPreview[]));
    const selectValue = ref<string>(storeToUpdate.value?.establishment_id ?? "");

    const schema = z.object({
        id: z.string(),
        name: z.string().min(3, "Campo obrigatório"),
        legal_name: z.string().min(3, "Campo obrigatório"),
        number: z.string().min(1, "Campo obrigatório"),

        establishment_id: z.string(),
        
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
        establishment_id:'',
        street:'',
        address_number:'',
        city:'',
        state:'',
        zip_code:'',
    });

    watch(storeToUpdate, (newStore: any) => {
        if (newStore) {
            state.id = newStore.id;
            state.name = newStore.name;
            state.legal_name = newStore.legal_name;
            state.number = newStore.number;
            state.establishment_id = newStore.establishment_id;
            state.street = newStore.street;
            state.address_number = newStore.number;
            state.city = newStore.city;
            state.state = newStore.state;
            state.zip_code = newStore.zip_code;

            selectValue.value = newStore.establishment_id;
        }
    });

    const openCreating = () => {
        Object.assign(state, {
            id: '',
            name: '',
            legal_name: '',
            number: '',
            establishment_id: '',
            street: '',
            address_number: '',
            city: '',
            state: '',
            zip_code: '',
        });

        selectValue.value = '';
        creating.value = true;
    };

    const closeCreating = () => {
        creating.value = false;
    };

    const openModal = async (store: StorePreview) => {
        storeToUpdate.value = store;
        isModalOpen.value = true;
    }

    const closeModal = async () => {
        isModalOpen.value = false;
    }

    function mapStoreData(stores: StorePreview[]) {
        return stores.map((store: StorePreview): DataPreview => ({
            id: store.id,
            name: store.name,
            legal_name: store.legal_name,
            number: store.number,
            establishment_id: store.establishment_id,
            street: store.address.street,
            address_number: store.address.number,
            city: store.address.city,
            state: store.address.state,
            zip_code: store.address.zip_code,
        }));
    }


    const createStore = async (form: { data: Schema }) => {
        const values = form.data;
        const payload = {
            name: values.name,
            legal_name: values.legal_name,
            number: values.number,
            establishment_id: values.establishment_id,
            address: {
                street: values.street,
                number: values.address_number,
                city: values.city,
                state: values.state,
                zip_code: values.zip_code,
            }
        };

        try {
            await create(payload);

            const [freshStores, freshEstabs] = await Promise.all([
                list(),
                listEstablishments(),
            ]);

            data.value = mapStoreData(freshStores);
            selectItems.value = mapSelectItems(freshEstabs);

            toast.add({ title: 'Loja criada com sucesso!', color: 'success' });
            creating.value = false;
        } catch (err) {
            toast.add({ title: 'Erro ao criar loja', color: 'error' });
            console.error(err);
        }
    };


    const updateStore = async (store: Partial<Schema>) => {
        const values = store.data;
        try {
            const updatePayload = {
                name: values.name,
                legal_name: values.legal_name,
                number: values.number,
                establishment_id: values.establishment_id,
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

            await update(values.id, updatePayload as StorePreview);

            const freshList = await list();
            data.value = mapStoreData(freshList);

            toast.add({ title: 'Estabelecimento atualizado com sucesso!', color: 'success' });
            isModalOpen.value = false;
        } catch (err) {
            toast.add({ title: 'Erro ao atualizar estabelecimento', color: 'error' });
            console.error(err);
        }
    }


    const removeAndRefresh = async (id: string) => {
        try {
            await remove(id);
            const freshList = await list();
            data.value = mapStoreData(freshList);
            toast.add({ title: 'Store deleted.', color: 'success' });
        } catch (err) {
            toast.add({ title: 'Erro ao deletar', color: 'error' });
            console.error(err);
        }
    };

    type DataPreview = {
        id: string;
        name: string;
        legal_name: string;
        number: string;
        establishment_id: string;
        street: string;
        address_number: string;
        city: string;
        state: string;
        zip_code: string;
    }

    const storesPreview = mapStoreData(stores.value as StorePreview[]);

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
            accessorKey: 'establishment_id',
            header: ({ column }) => {
                const isSorted = column.getIsSorted();
                return h(NuxtButton, {
                    color: 'neutral',
                    variant: 'ghost',
                    label: 'Estabelecimento',
                    icon: isSorted
                    ? isSorted === 'asc'
                        ? 'i-lucide-arrow-up-narrow-wide'
                        : 'i-lucide-arrow-down-wide-narrow'
                    : 'i-lucide-arrow-up-down',
                    class: '-mx-2.5 text-right',
                    onClick: () => column.toggleSorting(column.getIsSorted() === 'asc')
                })
            },
            cell: ({ row }) => {
                const estab = establishments.value?.find(est => est.id === row.getValue('establishment_id'));
                const name = estab?.name ?? '—';
                return h('div', { class: 'text-left font-medium' }, name);
            },
        },
        {
            accessorKey: 'street',
            header: () => h('fiv', { class: 'text-right' }, 'Rua'),
            cell: ({ row }) => h('div', { class: 'text-left font-medium' }, row.getValue('street')),
        },
        {
            accessorKey: 'address_number',
            header: () => h('fiv', { class: 'text-right' }, 'Numero Endereço'),
            cell: ({ row }) => h('div', { class: 'text-left font-medium' }, row.getValue('address_number')),
        },
        {
            accessorKey: 'city',
            header: () => h('fiv', { class: 'text-right' }, 'Cidade'),
            cell: ({ row }) => h('div', { class: 'text-left font-medium' }, row.getValue('city')),
        },
        {
            accessorKey: 'state',
            header: () => h('fiv', { class: 'text-right' }, 'Estado'),
            cell: ({ row }) => h('div', { class: 'text-left font-medium' }, row.getValue( 'state')),
        },
        {
            accessorKey: 'zip_code',
            header: () => h('fiv', { class: 'text-right' }, 'CEP'),
            cell: ({ row }) => h('div', { class: 'text-left font-medium' }, row.getValue('zip_code')),
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
                label: 'Copy store ID',
                onSelect() {
                    copy(row.original.id);
                    toast.add({
                        title: 'Store ID copied to clipboard!',
                        color: 'success',
                        icon: 'i-lucide-circle-check',
                    });
                }
            },
            {
                type: 'separator',
            },
            {
                label: 'View store details',
                onSelect() {
                    openModal(row.original as any);
                }
            },
            {
                label: 'Delete store',
                onSelect() {
                    removeAndRefresh(row.original.id);
                }
            }
        ];
    }

    const data = ref<DataPreview[]>(storesPreview);
</script>
