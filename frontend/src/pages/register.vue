<script setup lang="ts">
    import * as z from 'zod/v4';

    useHead({
        title: 'Register Page',
    });

    const router = useRouter();
    const auth = useAuth();
    const toast = useToast();

    const schema = z.object({
        name: z.string(),
        email: z.string().email('Invalid email'),
        password: z.string().min(3, 'Must be at least 3 characters')
    });

    type Schema = z.output<typeof schema>;
    type Register = {name: string, email: string, password: string};

    const state = reactive<Partial<Schema>>({
        name: undefined,
        email: undefined,
        password: undefined,
    });

    const register = async ({name, email, password}: Register) => {
        const success = await auth.register(name, email, password);
        if (success) {
            router.push('/login');
        } else {
            alert('Email inválido');
        }
    }

    async function onSubmit() {
        await register(state as Register);
        toast.success({ title: 'Success', message: 'The form has been submitted.', timeout: 500 });
    }
</script>


<!-- TEMPLATE -->

<template>
    <div class="flex flex-col pb-4 h-[100vh] items-center justify-between">        
        <div class="flex mt-8">
            <Icon name="svg:logo-white" class="w-16 h-16" />
        </div>
        
        <NuxtForm :schema="schema" :state="state" @submit="onSubmit"
            class="flex flex-col w-full mb-8 px-8 py-12 gap-6 rounded-2xl sm:w-md sm:px-14 sm:py-10 sm:bg-[#ffffff19]"
        >
            <p class="text-3xl text-center text-white">Acesse sua conta</p>

            <div class="flex flex-col gap-8">
                <div class="flex flex-col gap-2">
                    <label class="block text-sm font-bold text-gray-900 dark:text-white">Name</label>
                    <input
                        required
                        v-model="state.name"
                        name="name"
                        type="text"
                        placeholder="Digite seu nome"
                        class="sm:w-full sm:bg-background text-white p-3 rounded-xl bg-cyan-800 border-cyan-950 border-2 focus:outline-0 focus:border-white"
                    />
                </div>

                <div class="flex flex-col gap-2">
                    <label class="block text-sm font-bold text-gray-900 dark:text-white">Email</label>
                    <input
                        required
                        v-model="state.email"
                        name="email"
                        type="email"
                        placeholder="example@company.com"
                        class="sm:w-full sm:bg-background text-white p-3 rounded-xl bg-cyan-800 border-cyan-950 border-2 focus:outline-0 focus:border-white"
                    />
                </div>

                <div class="flex flex-col gap-2">
                    <label class="block text-sm font-bold text-gray-900 dark:text-white">Password</label>
                    <input
                        required
                        v-model="state.password"
                        name="password"
                        type="password"
                        placeholder="Sua senha"
                        class="sm:w-full sm:bg-background text-white p-3 rounded-xl bg-cyan-800 border-cyan-950 border-2 focus:outline-0 focus:border-white"
                    />
                </div>
            </div>

            <button type="submit" class="w-full p-3 mt-4 rounded-lg cursor-pointer text-sm font-bold bg-[#e0e0e0] border-[#e0e0e0] text-[#767676] hover:bg-[#f1f1f1]">
                Criar conta
            </button>

            <NuxtLink href="login" class="text-white text-center underline">
                Já tenho uma conta
            </NuxtLink>

            <hr class="my-2 h-0.5 border-t-0 bg-neutral-100 dark:bg-white" />

            <div class="flex items-center justify-center">
                <button class="flex w-64 gap-2 p-2 justify-evenly cursor-pointer text-sm font-bold bg-[#e0e0e0] border-[#e0e0e0] text-[#767676] hover:bg-[#f1f1f1]">
                    <Icon name="logos:google-icon" class="w-5 h-5" />
                    Acessar conta com Google
                </button>
            </div>
        </NuxtForm>

        <div>
            <p class="flex text-sm">
                Este site é protegido por reCAPTCHA e o Google.
                São aplicados os&nbsp;<a href="#" class="underline">Termos de serviço</a>&nbsp;e a&nbsp;<a href="#" class="underline">Política de privacidade</a>.
            </p>
        </div>
    </div>
</template>
