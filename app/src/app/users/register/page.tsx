"use client";

import FileUploadButton from "@/components/FileUploadButton";
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
import { cn } from "@/lib/utils";
import { createUser } from "@/services/user";
import { zodResolver } from "@hookform/resolvers/zod";
import { SubmitHandler, useForm } from "react-hook-form";

export default function Page() {
  const form = useForm<User>({
    resolver: zodResolver(UserSchema),
  });

  function handleProfileImageOnChange(value: string) {
    form.setValue("profileImage", value);
  }

  const onSubmit: SubmitHandler<User> = async (values) => {
    await createUser(values);
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <div className="mt-10 space-y-8 border-b border-gray-900/10 pb-12 sm:space-y-0 sm:divide-y sm:divide-gray-900/10 sm:border-t sm:pb-0">
          <FormField
            control={form.control}
            name="profileImage"
            defaultValue=""
            render={({ field }) => (
              <FormItem className="sm:grid sm:grid-cols-3 sm:items-start sm:gap-4 sm:py-6">
                <FormLabel className="user-form-label">Profile Image</FormLabel>
                <div className="mt-2 sm:col-span-2 sm:mt-0">
                  <div className="flex items-center space-x-2">
                    <div
                      className={cn(
                        "user-form-input-div",
                        "ring-0 ring-offset-0 shadow-none",
                        "focus-within:ring-0 focus:ring-offset-0"
                      )}
                    >
                      {form.watch("username") !== "" ? (
                        <FileUploadButton
                          filename={`profile-image-${form.watch("username")}`}
                          saveDir="user-profile-images"
                          setValue={handleProfileImageOnChange}
                        />
                      ) : (
                        <div className="flex items-center space-x-2">
                          <span className="text-gray-500">
                            Please input username first
                          </span>
                        </div>
                      )}
                    </div>
                    <FormControl>
                      <Input {...field} disabled className="border-none" />
                    </FormControl>
                  </div>

                  <FormDescription>
                    This is your public display profile image.
                  </FormDescription>
                  <FormMessage />
                </div>
              </FormItem>
            )}
          />

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
            name="position"
            defaultValue=""
            render={({ field }) => (
              <FormItem className="sm:grid sm:grid-cols-3 sm:items-start sm:gap-4 sm:py-6">
                <FormLabel className="user-form-label">Position</FormLabel>
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
                  <div
                    className={cn(
                      "user-form-input-div",
                      "ring-0 ring-offset-0 shadow-none",
                      "focus-within:ring-0 focus:ring-offset-0"
                    )}
                  >
                    <div className="flex items-center space-x-2">
                      <FormControl>
                        <Switch
                          checked={field.value}
                          onCheckedChange={field.onChange}
                        />
                      </FormControl>
                      <span>{field.value ? "Active" : "Non Active"}</span>
                    </div>
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
