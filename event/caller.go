// Code generated by /tools/cmd/generatecmdcaller.go; DO NOT EDIT.
//go:generate go run ../tools/cmd/generatecmdcaller.go

package event

import (
    "fmt"
    "github.com/TicketsBot/worker/bot/command"
    cmdcontext "github.com/TicketsBot/worker/bot/command/context"
    "github.com/TicketsBot/worker/bot/command/impl/general"
    "github.com/TicketsBot/worker/bot/command/impl/settings"
    "github.com/TicketsBot/worker/bot/command/impl/tickets"
    "github.com/TicketsBot/worker/bot/command/impl/admin"
    "github.com/TicketsBot/worker/bot/command/impl/statistics"
    "github.com/TicketsBot/worker/bot/command/impl/settings/setup"
    "github.com/TicketsBot/worker/bot/command/impl/tags"
    "github.com/TicketsBot/worker/bot/command/registry"
    "github.com/pkg/errors"
    "github.com/rxdn/gdl/objects/interaction"
    "strconv"
)

var ErrArgumentNotFound = errors.New("argument not found")

func callCommand(
    cmd registry.Command,
    ctx *cmdcontext.SlashCommandContext,
    options []interaction.ApplicationCommandInteractionDataOption,
) error {
    switch v := cmd.(type) {
    
    case admin.AdminBlacklistCommand:
        var arg0 string

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt0.Name)
            }
            arg0 = argValue
        }
        var arg1 *string

        opt1, ok1 := findOption(cmd.Properties().Arguments[1], options)
        if !ok1 {
            arg1 = nil
        } else { 
            argValue, ok := opt1.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt1.Name)
            }
            arg1 = &argValue
        }

        v.Execute(ctx, arg0, arg1)
    case admin.AdminCheckBlacklistCommand:
        var arg0 string

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case admin.AdminCheckPremiumCommand:
        var arg0 string

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case admin.AdminCommand:

        v.Execute(ctx)
    case admin.AdminGenPremiumCommand:
        var arg0 string

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt0.Name)
            }
            arg0 = argValue
        }
        var arg1 int

        opt1, ok1 := findOption(cmd.Properties().Arguments[1], options)
        if !ok1 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt1.Value.(float64)
            if !ok {
                return fmt.Errorf("option %s was not a float64", opt1.Name)
            }
            arg1 = int(argValue)
        }
        var arg2 *int

        opt2, ok2 := findOption(cmd.Properties().Arguments[2], options)
        if !ok2 {
            arg2 = nil
        } else { 
            argValue, ok := opt2.Value.(float64)
            if !ok {
                return fmt.Errorf("option %s was not a float64", opt2.Name)
            }
            tmp := int(argValue)
            arg2 = &tmp
        }

        v.Execute(ctx, arg0, arg1, arg2)
    case admin.AdminGetOwnerCommand:
        var arg0 string

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case admin.AdminListGuildEntitlementsCommand:
        var arg0 string

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case admin.AdminListUserEntitlementsCommand:
        var arg0 uint64

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else {
            raw, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a snowflake", opt0.Name)
            }

            argValue, err := strconv.ParseUint(raw, 10, 64)
            if err != nil {
                return fmt.Errorf("option %s was not a valid snowflake", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case admin.AdminRecacheCommand:
        var arg0 *string

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            arg0 = nil
        } else { 
            argValue, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt0.Name)
            }
            arg0 = &argValue
        }

        v.Execute(ctx, arg0)
    case admin.AdminWhitelabelAssignGuildCommand:
        var arg0 string

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt0.Name)
            }
            arg0 = argValue
        }
        var arg1 string

        opt1, ok1 := findOption(cmd.Properties().Arguments[1], options)
        if !ok1 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt1.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt1.Name)
            }
            arg1 = argValue
        }

        v.Execute(ctx, arg0, arg1)
    case admin.AdminWhitelabelDataCommand:
        var arg0 uint64

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else {
            raw, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a snowflake", opt0.Name)
            }

            argValue, err := strconv.ParseUint(raw, 10, 64)
            if err != nil {
                return fmt.Errorf("option %s was not a valid snowflake", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case general.AboutCommand:

        v.Execute(ctx)
    case general.HelpCommand:

        v.Execute(ctx)
    case general.InviteCommand:

        v.Execute(ctx)
    case general.JumpToTopCommand:

        v.Execute(ctx)
    case general.VoteCommand:

        v.Execute(ctx)
    case settings.AddAdminCommand:
        var arg0 uint64

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else {
            raw, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a snowflake", opt0.Name)
            }

            argValue, err := strconv.ParseUint(raw, 10, 64)
            if err != nil {
                return fmt.Errorf("option %s was not a valid snowflake", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case settings.AddSupportCommand:
        var arg0 uint64

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else {
            raw, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a snowflake", opt0.Name)
            }

            argValue, err := strconv.ParseUint(raw, 10, 64)
            if err != nil {
                return fmt.Errorf("option %s was not a valid snowflake", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case settings.AutoCloseCommand:

        v.Execute(ctx)
    case settings.AutoCloseConfigureCommand:

        v.Execute(ctx)
    case settings.AutoCloseExcludeCommand:

        v.Execute(ctx)
    case settings.BlacklistCommand:
        var arg0 uint64

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else {
            raw, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a snowflake", opt0.Name)
            }

            argValue, err := strconv.ParseUint(raw, 10, 64)
            if err != nil {
                return fmt.Errorf("option %s was not a valid snowflake", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case settings.LanguageCommand:

        v.Execute(ctx)
    case settings.PanelCommand:

        v.Execute(ctx)
    case settings.PremiumCommand:

        v.Execute(ctx)
    case settings.RemoveAdminCommand:
        var arg0 uint64

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else {
            raw, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a snowflake", opt0.Name)
            }

            argValue, err := strconv.ParseUint(raw, 10, 64)
            if err != nil {
                return fmt.Errorf("option %s was not a valid snowflake", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case settings.RemoveSupportCommand:
        var arg0 uint64

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else {
            raw, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a snowflake", opt0.Name)
            }

            argValue, err := strconv.ParseUint(raw, 10, 64)
            if err != nil {
                return fmt.Errorf("option %s was not a valid snowflake", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case settings.ViewStaffCommand:

        v.Execute(ctx)
    case setup.AutoSetupCommand:

        v.Execute(ctx)
    case setup.LimitSetupCommand:
        var arg0 int

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt0.Value.(float64)
            if !ok {
                return fmt.Errorf("option %s was not a float64", opt0.Name)
            }
            arg0 = int(argValue)
        }

        v.Execute(ctx, arg0)
    case setup.SetupCommand:

        v.Execute(ctx)
    case setup.ThreadsSetupCommand:
        var arg0 bool

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt0.Value.(bool)
            if !ok {
                return fmt.Errorf("option %s was not a bool", opt0.Name)
            }
            arg0 = argValue

            
        }
        var arg1 *uint64

        opt1, ok1 := findOption(cmd.Properties().Arguments[1], options)
        if !ok1 {
            arg1 = nil
        } else {
            raw, ok := opt1.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a snowflake", opt1.Name)
            }

            argValue, err := strconv.ParseUint(raw, 10, 64)
            if err != nil {
                return fmt.Errorf("option %s was not a valid snowflake", opt1.Name)
            }
            arg1 = &argValue
        }

        v.Execute(ctx, arg0, arg1)
    case setup.TranscriptsSetupCommand:
        var arg0 uint64

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else {
            raw, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a snowflake", opt0.Name)
            }

            argValue, err := strconv.ParseUint(raw, 10, 64)
            if err != nil {
                return fmt.Errorf("option %s was not a valid snowflake", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case statistics.StatsCommand:

        v.Execute(ctx)
    case statistics.StatsServerCommand:

        v.Execute(ctx)
    case statistics.StatsUserCommand:
        var arg0 uint64

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else {
            raw, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a snowflake", opt0.Name)
            }

            argValue, err := strconv.ParseUint(raw, 10, 64)
            if err != nil {
                return fmt.Errorf("option %s was not a valid snowflake", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case tags.ManageTagsAddCommand:
        var arg0 string

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt0.Name)
            }
            arg0 = argValue
        }
        var arg1 string

        opt1, ok1 := findOption(cmd.Properties().Arguments[1], options)
        if !ok1 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt1.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt1.Name)
            }
            arg1 = argValue
        }

        v.Execute(ctx, arg0, arg1)
    case tags.ManageTagsCommand:

        v.Execute(ctx)
    case tags.ManageTagsDeleteCommand:
        var arg0 string

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case tags.ManageTagsListCommand:

        v.Execute(ctx)
    case tags.TagCommand:
        var arg0 string

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case tickets.AddCommand:
        var arg0 uint64

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else {
            raw, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a snowflake", opt0.Name)
            }

            argValue, err := strconv.ParseUint(raw, 10, 64)
            if err != nil {
                return fmt.Errorf("option %s was not a valid snowflake", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case tickets.ClaimCommand:

        v.Execute(ctx)
    case tickets.CloseCommand:
        var arg0 *string

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            arg0 = nil
        } else { 
            argValue, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt0.Name)
            }
            arg0 = &argValue
        }

        v.Execute(ctx, arg0)
    case tickets.CloseRequestCommand:
        var arg0 *int

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            arg0 = nil
        } else { 
            argValue, ok := opt0.Value.(float64)
            if !ok {
                return fmt.Errorf("option %s was not a float64", opt0.Name)
            }
            tmp := int(argValue)
            arg0 = &tmp
        }
        var arg1 *string

        opt1, ok1 := findOption(cmd.Properties().Arguments[1], options)
        if !ok1 {
            arg1 = nil
        } else { 
            argValue, ok := opt1.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt1.Name)
            }
            arg1 = &argValue
        }

        v.Execute(ctx, arg0, arg1)
    case tickets.NotesCommand:

        v.Execute(ctx)
    case tickets.OnCallCommand:

        v.Execute(ctx)
    case tickets.OpenCommand:
        var arg0 *string

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            arg0 = nil
        } else { 
            argValue, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt0.Name)
            }
            arg0 = &argValue
        }

        v.Execute(ctx, arg0)
    case tickets.RemoveCommand:
        var arg0 uint64

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else {
            raw, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a snowflake", opt0.Name)
            }

            argValue, err := strconv.ParseUint(raw, 10, 64)
            if err != nil {
                return fmt.Errorf("option %s was not a valid snowflake", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case tickets.RenameCommand:
        var arg0 string

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a string", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case tickets.ReopenCommand:
        var arg0 int

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt0.Value.(float64)
            if !ok {
                return fmt.Errorf("option %s was not a float64", opt0.Name)
            }
            arg0 = int(argValue)
        }

        v.Execute(ctx, arg0)
    case tickets.StartTicketCommand:

        v.Execute(ctx)
    case tickets.SwitchPanelCommand:
        var arg0 int

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else { 
            argValue, ok := opt0.Value.(float64)
            if !ok {
                return fmt.Errorf("option %s was not a float64", opt0.Name)
            }
            arg0 = int(argValue)
        }

        v.Execute(ctx, arg0)
    case tickets.TransferCommand:
        var arg0 uint64

        opt0, ok0 := findOption(cmd.Properties().Arguments[0], options)
        if !ok0 {
            return ErrArgumentNotFound
        } else {
            raw, ok := opt0.Value.(string)
            if !ok {
                return fmt.Errorf("option %s was not a snowflake", opt0.Name)
            }

            argValue, err := strconv.ParseUint(raw, 10, 64)
            if err != nil {
                return fmt.Errorf("option %s was not a valid snowflake", opt0.Name)
            }
            arg0 = argValue
        }

        v.Execute(ctx, arg0)
    case tickets.UnclaimCommand:

        v.Execute(ctx)

    
    case tags.TagAliasCommand:
        v.Execute(ctx)

    default:
        return fmt.Errorf("unknown command %s", cmd.Properties().Name)
    }

    return nil
}

func findOption(
    arg command.Argument,
    options []interaction.ApplicationCommandInteractionDataOption,
) (interaction.ApplicationCommandInteractionDataOption, bool) {
    for _, option := range options {
        if option.Name == arg.Name {
            return option, true
        }
    }

    return interaction.ApplicationCommandInteractionDataOption{}, false
}
