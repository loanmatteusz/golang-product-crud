<template>
    <div class="flex flex-col p-4 items-center">
        <NuxtForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit">
            <div class="flex flex-col p-8 rounded-md gap-4 items-center">
                <NuxtFormField label="Email" name="email">
                  <NuxtInput v-model="state.email" />
                </NuxtFormField>
            
                <NuxtFormField label="Password" name="password">
                  <NuxtInput v-model="state.password" type="password" />
                </NuxtFormField>
            
                <NuxtButton type="submit">Submit</NuxtButton>

                <NuxtLink to="/register" class="text-sm text-blue-500 hover:underline mt-2">
                    Criar uma conta
                </NuxtLink>
            </div>
        </NuxtForm>
    </div>
</template>


<script setup lang="ts">
    import * as z from 'zod'
    import type { FormSubmitEvent } from '@nuxt/ui'
    
    const router = useRouter();
    const auth = useAuth();
    const toast = useToast();

    const schema = z.object({
        email: z.string().email('Invalid email'),
        password: z.string().min(3, 'Must be at least 3 characters')
    });

    type Schema = z.output<typeof schema>;

    const state = reactive<Partial<Schema>>({
        email: undefined,
        password: undefined
    });

    const login = async ({email, password}: {email: string, password: string}) => {
        const success = await auth.login(email, password);
        if (success) {
            router.push('/');
        } else {
            alert('Credenciais inválidas');
        }
    }

    async function onSubmit(event: FormSubmitEvent<Schema>) {
        await login(state);
        toast.add({ title: 'Success', description: 'The form has been submitted.', color: 'success' });
    }
</script>
