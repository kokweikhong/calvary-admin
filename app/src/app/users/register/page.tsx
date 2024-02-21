"use client";

import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Switch } from "@/components/ui/switch";
import { UserSchema, type User } from "@/interfaces/user";
import { zodResolver } from "@hookform/resolvers/zod";
import { SubmitHandler, useForm } from "react-hook-form";

export default function Page() {
  const form = useForm<User>({
    resolver: zodResolver(UserSchema),
  });

  const onSubmit: SubmitHandler<User> = async (values) => {
    console.log(values);
  };

  console.log(form.formState.errors);

  // id: z.number(),
  // username: z.string(),
  // email: z.string(),
  // password: z.string(),
  // role: z.string(),
  // department: z.string(),
  // profileImage: z.string(),
  // isActive: z.boolean(),
  // position: z.string(),
  // updatedAt: z.string(),
  // createdAt: z.string(),

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <div className="mt-10 space-y-8 border-b border-gray-900/10 pb-12 sm:space-y-0 sm:divide-y sm:divide-gray-900/10 sm:border-t sm:pb-0">
          <FormField
            control={form.control}
            name="username"
            defaultValue=""
            render={({ field }) => (
              <FormItem className="sm:grid sm:grid-cols-3 sm:items-start sm:gap-4 sm:py-6">
                <FormLabel className="user-form-label">Username</FormLabel>
                <div className="mt-2 sm:col-span-2 sm:mt-0">
                  <div className="user-form-input-div">
                    <FormControl>
                      <Input placeholder="your unique username" {...field} />
                    </FormControl>
                  </div>
                  <FormDescription>
                    This is your public display name.
                  </FormDescription>
                  <FormMessage />
                </div>
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="email"
            defaultValue=""
            render={({ field }) => (
              <FormItem className="sm:grid sm:grid-cols-3 sm:items-start sm:gap-4 sm:py-6">
                <FormLabel className="user-form-label">Email Address</FormLabel>
                <div className="mt-2 sm:col-span-2 sm:mt-0">
                  <div className="user-form-input-div">
                    <FormControl>
                      <Input type="email" autoComplete="email" {...field} />
                    </FormControl>
                  </div>
                  <FormDescription>
                    This is your email address with @calvarycarpentry.com
                  </FormDescription>
                  <FormMessage />
                </div>
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="password"
            defaultValue=""
            render={({ field }) => (
              <FormItem className="sm:grid sm:grid-cols-3 sm:items-start sm:gap-4 sm:py-6">
                <FormLabel className="user-form-label">Password</FormLabel>
                <div className="mt-2 sm:col-span-2 sm:mt-0">
                  <div className="user-form-input-div">
                    <FormControl>
                      <Input type="password" {...field} />
                    </FormControl>
                  </div>
                  <FormDescription>
                    This is your secret password.
                  </FormDescription>
                  <FormMessage />
                </div>
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="role"
            defaultValue=""
            render={({ field }) => (
              <FormItem className="sm:grid sm:grid-cols-3 sm:items-start sm:gap-4 sm:py-6">
                <FormLabel className="user-form-label">Role</FormLabel>
                <div className="mt-2 sm:col-span-2 sm:mt-0">
                  <div className="user-form-input-div">
                    <Select
                      onValueChange={field.onChange}
                      defaultValue={field.value}
                    >
                      <FormControl>
                        <SelectTrigger>
                          <SelectValue placeholder="Select a role" />
                        </SelectTrigger>
                      </FormControl>
                      <SelectContent>
                        <SelectItem value="m@example.com">
                          m@example.com
                        </SelectItem>
                        <SelectItem value="m@google.com">
                          m@google.com
                        </SelectItem>
                        <SelectItem value="m@support.com">
                          m@support.com
                        </SelectItem>
                      </SelectContent>
                    </Select>
                  </div>
                  <FormDescription>
                    This is the role for this admin system.
                  </FormDescription>
                  <FormMessage />
                </div>
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="department"
            defaultValue=""
            render={({ field }) => (
              <FormItem className="sm:grid sm:grid-cols-3 sm:items-start sm:gap-4 sm:py-6">
                <FormLabel className="user-form-label">Department</FormLabel>
                <div className="mt-2 sm:col-span-2 sm:mt-0">
                  <div className="user-form-input-div">
                    <Select
                      onValueChange={field.onChange}
                      defaultValue={field.value}
                    >
                      <FormControl>
                        <SelectTrigger>
                          <SelectValue placeholder="Select the designated department" />
                        </SelectTrigger>
                      </FormControl>
                      <SelectContent>
                        <SelectItem value="m@example.com">
                          m@example.com
                        </SelectItem>
                        <SelectItem value="m@google.com">
                          m@google.com
                        </SelectItem>
                        <SelectItem value="m@support.com">
                          m@support.com
                        </SelectItem>
                      </SelectContent>
                    </Select>
                  </div>
                  <FormDescription>
                    This is the role for this admin system.
                  </FormDescription>
                  <FormMessage />
                </div>
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="isActive"
            defaultValue={true}
            render={({ field }) => (
              <FormItem className="sm:grid sm:grid-cols-3 sm:items-start sm:gap-4 sm:py-6">
                <FormLabel className="user-form-label">isActive</FormLabel>
                <div className="mt-2 sm:col-span-2 sm:mt-0">
                  <div className="user-form-input-div">
                    <FormControl>
                      <Switch
                        checked={field.value}
                        onCheckedChange={field.onChange}
                      />
                    </FormControl>
                  </div>
                  <FormDescription>
                    This user is active in the system.
                  </FormDescription>
                  <FormMessage />
                </div>
              </FormItem>
            )}
          />

          <Button type="submit">Submit</Button>
        </div>
      </form>
    </Form>
  );
}
